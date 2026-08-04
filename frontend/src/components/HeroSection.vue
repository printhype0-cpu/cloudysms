<template>
  <div class="relative flex flex-col items-center justify-center py-16 md:py-24 px-4 sm:px-6 text-center mt-8 md:mt-12 mb-16 md:mb-24 overflow-hidden">
    <!-- Animated background glow -->
    <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[120%] max-w-[800px] h-[400px] bg-gradient-to-r from-emerald-500/20 to-teal-500/20 blur-[100px] rounded-full -z-10 animate-pulse" style="animation-duration: 4s;"></div>

    <!-- Top Badge -->
    <div class="animate-list-in inline-flex items-center gap-1.5 sm:gap-2 px-3 sm:px-4 py-1.5 rounded-full bg-slate-900/40 border border-white/10 text-emerald-400 text-xs sm:text-sm font-medium mb-6 sm:mb-8 flex-wrap justify-center backdrop-blur-sm" style="animation-delay: 0.1s;">
      <span>✨</span>
      <span class="text-white/40 hidden sm:inline">—</span>
      <span class="text-center">AI-Powered WhatsApp Marketing Platform</span>
    </div>

    <!-- Main Heading -->
    <h1 class="animate-list-in text-4xl sm:text-5xl md:text-6xl lg:text-[5rem] leading-[1.1] md:leading-tight font-extrabold tracking-tight mb-6 w-full break-words" style="animation-delay: 0.2s;">
      <span class="text-gradient-primary block mb-2">Unlock Explosive Growth.</span>
      <span class="text-[color:var(--text-main)] block">Turn WhatsApp into Your #1 Sales Channel</span>
    </h1>

    <!-- Subheading -->
    <p class="animate-list-in text-base sm:text-lg md:text-xl text-[color:var(--text-muted)] max-w-3xl mx-auto mb-8 sm:mb-10 leading-relaxed px-2" style="animation-delay: 0.3s;">
      Broadcast, Automate, Engage, Sell - do everything with the AI-powered <br class="hidden md:block"/>
      WhatsApp Marketing & Engagement Platform. Powered by Official WhatsApp APIs ⚡
    </p>

    <!-- Buttons -->
    <div class="animate-list-in flex flex-col sm:flex-row items-center justify-center gap-3 sm:gap-4 w-full max-w-md sm:max-w-none mx-auto" style="animation-delay: 0.4s;">
      <button @click="router.push('/login')" class="group w-full sm:w-auto px-6 sm:px-8 py-3.5 sm:py-4 bg-emerald-600 hover:bg-emerald-500 text-white font-bold rounded-xl transition-all shadow-[0_0_20px_rgba(16,185,129,0.3)] hover:shadow-[0_0_30px_rgba(16,185,129,0.5)] hover:-translate-y-1 flex items-center justify-center gap-2 text-base sm:text-lg">
        Start 14-Day FREE Trial
        <svg xmlns="http://www.w3.org/2000/svg" class="transition-transform group-hover:translate-x-1" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M5 12h14"></path>
          <path d="m12 5 7 7-7 7"></path>
        </svg>
      </button>
      
      <button @click="router.push('/contact')" class="w-full sm:w-auto px-6 sm:px-8 py-3.5 sm:py-4 bg-white/5 hover:bg-white/10 text-white font-bold rounded-xl border border-white/10 transition-all hover:-translate-y-1 flex items-center justify-center gap-2 text-base sm:text-lg backdrop-blur-sm">
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
