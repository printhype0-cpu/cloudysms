<template>
  <div class="flex flex-col items-center justify-center py-12 md:py-20 px-4 sm:px-6 text-center mt-8 md:mt-12 mb-16 md:mb-24 overflow-hidden">
    <!-- Top Badge -->
    <div class="inline-flex items-center gap-1.5 sm:gap-2 px-3 sm:px-4 py-1.5 rounded-full bg-purple-50 border border-purple-100 text-purple-600 text-xs sm:text-sm font-medium mb-6 sm:mb-8 flex-wrap justify-center">
      <span>✨</span>
      <span class="text-purple-300 hidden sm:inline">—</span>
      <span class="text-center">AI-Powered WhatsApp Marketing Platform</span>
    </div>

    <!-- Main Heading -->
    <h1 class="text-4xl sm:text-5xl md:text-6xl lg:text-[5rem] leading-[1.1] md:leading-tight font-extrabold tracking-tight mb-6 w-full break-words">
      <span class="text-[#00d26a] block mb-2">Unlock Explosive Growth.</span>
      <span class="text-[color:var(--text-main)] block">Turn WhatsApp into Your #1 Sales Channel</span>
    </h1>

    <!-- Subheading -->
    <p class="text-base sm:text-lg md:text-xl text-[color:var(--text-muted)] max-w-3xl mx-auto mb-8 sm:mb-10 leading-relaxed px-2">
      Broadcast, Automate, Engage, Sell - do everything with the AI-powered <br class="hidden md:block"/>
      WhatsApp Marketing & Engagement Platform. Powered by Official WhatsApp APIs ⚡
    </p>

    <!-- Buttons -->
    <div class="flex flex-col sm:flex-row items-center justify-center gap-3 sm:gap-4 w-full max-w-md sm:max-w-none mx-auto">
      <button @click="router.push('/login')" class="w-full sm:w-auto px-6 sm:px-8 py-3.5 sm:py-4 bg-[#00d26a] hover:bg-[#00c060] text-white font-bold rounded-lg transition-colors flex items-center justify-center gap-2 text-base sm:text-lg shadow-sm">
        Start 14-Day FREE Trial
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M5 12h14"></path>
          <path d="m12 5 7 7-7 7"></path>
        </svg>
      </button>
      
      <button @click="router.push('/contact')" class="w-full sm:w-auto px-6 sm:px-8 py-3.5 sm:py-4 bg-transparent hover:bg-slate-50/10 text-[color:var(--text-main)] font-bold rounded-lg border-2 border-[color:var(--text-main)] transition-colors flex items-center justify-center gap-2 text-base sm:text-lg">
        Join Live Demo
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M5 12h14"></path>
          <path d="m12 5 7 7-7 7"></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
</script>
