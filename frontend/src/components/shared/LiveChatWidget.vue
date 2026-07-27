<template>
  <div class="fixed bottom-6 right-6 z-50">
    <!-- Chat Window -->
    <Transition name="fade-scale">
      <div v-if="isOpen" class="absolute bottom-16 right-0 mb-4 w-[calc(100vw-3rem)] sm:w-96 max-h-[80vh] bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl shadow-2xl flex flex-col overflow-hidden transition-all duration-300 z-50">
        <!-- Header -->
        <div class="p-4 bg-gradient-to-r from-emerald-500 to-emerald-600 text-white flex justify-between items-center shrink-0">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-white/20 flex items-center justify-center">
              <Bot class="w-6 h-6 text-white" />
            </div>
            <div>
              <h3 class="font-semibold text-sm">Cloudy Bot</h3>
              <p class="text-xs text-emerald-100 flex items-center gap-1">
                <span class="w-2 h-2 rounded-full bg-green-400"></span> Online
              </p>
            </div>
          </div>
          <button @click="isOpen = false" class="p-1 hover:bg-white/20 rounded-lg transition-colors">
            <X class="w-5 h-5 text-white" />
          </button>
        </div>

        <!-- Messages Area -->
        <div class="flex-1 h-80 sm:h-96 overflow-y-auto p-4 bg-slate-50 dark:bg-[#0a0a0b] flex flex-col gap-3" ref="messagesContainer">
          <div v-for="(msg, index) in messages" :key="index" 
               class="max-w-[85%] rounded-2xl px-4 py-2.5 text-sm shadow-sm"
               :class="msg.role === 'user' 
                  ? 'bg-emerald-500 text-white self-end rounded-br-sm' 
                  : 'bg-white dark:bg-slate-800 text-slate-800 dark:text-slate-200 border border-slate-100 dark:border-slate-700 self-start rounded-bl-sm'">
            {{ msg.content }}
          </div>
          <div v-if="isTyping" class="bg-white dark:bg-slate-800 border border-slate-100 dark:border-slate-700 self-start rounded-2xl rounded-bl-sm px-4 py-3 shadow-sm flex items-center gap-1.5">
            <span class="w-1.5 h-1.5 rounded-full bg-slate-400 animate-bounce"></span>
            <span class="w-1.5 h-1.5 rounded-full bg-slate-400 animate-bounce" style="animation-delay: 0.15s"></span>
            <span class="w-1.5 h-1.5 rounded-full bg-slate-400 animate-bounce" style="animation-delay: 0.3s"></span>
          </div>
        </div>

        <!-- Input Area -->
        <div class="p-3 border-t border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 shrink-0">
          <form @submit.prevent="sendMessage" class="flex gap-2 items-center bg-slate-50 dark:bg-slate-800/50 rounded-xl p-1 border border-slate-200 dark:border-slate-700 focus-within:border-emerald-500 dark:focus-within:border-emerald-500 transition-colors">
            <input 
              v-model="inputValue" 
              type="text" 
              placeholder="Type your message..." 
              class="flex-1 bg-transparent border-none focus:outline-none px-3 py-2 text-sm text-slate-800 dark:text-slate-200"
            />
            <button 
              type="submit" 
              :disabled="!inputValue.trim()"
              class="p-2 rounded-lg bg-emerald-500 text-white disabled:opacity-50 disabled:cursor-not-allowed hover:bg-emerald-600 transition-colors"
            >
              <Send class="w-4 h-4" />
            </button>
          </form>
        </div>
      </div>
    </Transition>

    <!-- Floating Action Button -->
    <button 
      @click="isOpen = !isOpen"
      class="w-14 h-14 bg-emerald-500 hover:bg-emerald-600 text-white rounded-full flex items-center justify-center shadow-xl shadow-emerald-500/30 transition-transform hover:scale-105 active:scale-95"
      :class="isOpen ? 'scale-90' : ''"
    >
      <Transition name="fade" mode="out-in">
        <X v-if="isOpen" class="w-6 h-6" />
        <MessageCircle v-else class="w-6 h-6" />
      </Transition>
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue'
import { MessageCircle, X, Send, Bot } from 'lucide-vue-next'

const isOpen = ref(false)
const inputValue = ref('')
const isTyping = ref(false)
const messagesContainer = ref<HTMLElement | null>(null)

interface Message {
  role: 'bot' | 'user'
  content: string
}

const messages = ref<Message[]>([
  { role: 'bot', content: 'Hi there! 👋 Welcome to CloudySMS.' },
  { role: 'bot', content: 'How can I help you scale your WhatsApp marketing today?' }
])

const scrollToBottom = async () => {
  await nextTick()
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

const step = ref<'initial' | 'awaiting_email' | 'done'>('initial')
const userEmail = ref('')
const userMessage = ref('')

const sendMessage = async () => {
  if (!inputValue.value.trim()) return

  const messageText = inputValue.value.trim()
  inputValue.value = ''

  // Add user message to UI
  messages.value.push({ role: 'user', content: messageText })
  scrollToBottom()

  isTyping.value = true
  scrollToBottom()

  if (step.value === 'initial') {
    userMessage.value = messageText
    setTimeout(() => {
      isTyping.value = false
      messages.value.push({
        role: 'bot',
        content: "Thanks! What's the best email to reach you at? We'll get back to you shortly."
      })
      step.value = 'awaiting_email'
      scrollToBottom()
    }, 1000)
    return
  }

  if (step.value === 'awaiting_email') {
    userEmail.value = messageText
    
    // Validate email roughly
    if (!userEmail.value.includes('@') || !userEmail.value.includes('.')) {
      setTimeout(() => {
        isTyping.value = false
        messages.value.push({
          role: 'bot',
          content: "That doesn't look like a valid email. Could you try again?"
        })
        scrollToBottom()
      }, 500)
      return
    }

    step.value = 'done'

    try {
      // Send to real backend
      await fetch('/api/webhooks/website-lead', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          email: userEmail.value,
          message: userMessage.value
        })
      })

      setTimeout(() => {
        isTyping.value = false
        messages.value.push({
          role: 'bot',
          content: "Got it! We've saved your information and will be in touch soon. Have a great day!"
        })
        scrollToBottom()
      }, 1000)
    } catch (e) {
      setTimeout(() => {
        isTyping.value = false
        messages.value.push({
          role: 'bot',
          content: "Oops! Something went wrong on our end. Please try again later."
        })
        scrollToBottom()
      }, 1000)
    }
  } else if (step.value === 'done') {
      setTimeout(() => {
        isTyping.value = false
        messages.value.push({
          role: 'bot',
          content: "We've already received your message! We'll reply via email shortly."
        })
        scrollToBottom()
      }, 1000)
  }
}
</script>

<style scoped>
.fade-scale-enter-active,
.fade-scale-leave-active {
  transition: opacity 0.3s ease, transform 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  transform-origin: bottom right;
}

.fade-scale-enter-from,
.fade-scale-leave-to {
  opacity: 0;
  transform: scale(0.85) translateY(20px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: rotate(-90deg) scale(0.5);
}
</style>
