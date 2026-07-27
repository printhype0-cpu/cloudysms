<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'

useSeo({
  title: 'Changelog - CloudySMS',
  description: 'See the latest updates, new features, and bug fixes for the CloudySMS platform.'
})

const releases = [
  {
    version: 'v1.2.0',
    date: 'August 15, 2026',
    title: 'Visual Flow Builder & AI Contexts',
    content: 'We are thrilled to announce our brand new Visual Flow Builder! You can now design complex chatbot flows with a drag-and-drop interface. Additionally, we have added AI Contexts, allowing you to train your bots on your own company knowledge base.'
  },
  {
    version: 'v1.1.5',
    date: 'July 28, 2026',
    title: 'Advanced Analytics Dashboard',
    content: 'The analytics dashboard has been completely revamped. You can now track agent response times, message delivery rates, and customer satisfaction metrics with beautiful new visualizations.'
  },
  {
    version: 'v1.1.0',
    date: 'June 10, 2026',
    title: 'Team Inbox Enhancements',
    content: 'Collaborating on customer support just got easier. We have introduced internal notes, agent assignments, and collision detection to ensure your team never steps on each others toes.'
  }
]
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">What's <span class="text-gradient">new</span></h1>
        <p class="page-subtitle fade-in-up delay-1">The latest updates and improvements to CloudySMS.</p>
      </div>

      <div class="changelog-list fade-in-up delay-2">
        <div class="release-item" v-for="release in releases" :key="release.version">
          <div class="release-meta">
            <div class="version-badge">{{ release.version }}</div>
            <div class="release-date">{{ release.date }}</div>
          </div>
          <div class="release-content">
            <h3 class="release-title">{{ release.title }}</h3>
            <p class="release-body">{{ release.content }}</p>
          </div>
        </div>
      </div>
    </div>
  </MarketingLayout>
</template>

<style scoped>
.page-container {
  padding: 6rem 4rem;
  max-width: 900px;
  margin: 0 auto;
  width: 100%;
}

.page-header {
  text-align: left;
  margin-bottom: 5rem;
}

.page-title {
  font-size: 3.5rem;
  font-weight: 800;
  margin-bottom: 1rem;
  letter-spacing: -0.03em;
}

.text-gradient {
  color: var(--text-main);
}

.page-subtitle {
  font-size: 1.25rem;
  color: var(--text-muted);
}

.changelog-list {
  display: flex;
  flex-direction: column;
  gap: 4rem;
  position: relative;
}

.changelog-list::before {
  content: '';
  position: absolute;
  top: 0;
  bottom: 0;
  left: 120px;
  width: 2px;
  background: rgba(255, 255, 255, 0.1);
}

.release-item {
  display: flex;
  gap: 3rem;
  position: relative;
}

.release-meta {
  width: 120px;
  flex-shrink: 0;
  text-align: right;
  position: relative;
}

.release-meta::after {
  content: '';
  position: absolute;
  top: 0.5rem;
  right: -31px;
  width: 12px;
  height: 12px;
  background: #10b981;
  border: 3px solid #0a0a0b;
  border-radius: 50%;
  z-index: 2;
}

.version-badge {
  display: inline-block;
  background: rgba(16, 185, 129, 0.1);
  color: #10b981;
  border: 1px solid rgba(16, 185, 129, 0.2);
  padding: 0.25rem 0.75rem;
  border-radius: 1rem;
  font-size: 0.85rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.release-date {
  font-size: 0.9rem;
  color: var(--text-muted);
}

.release-content {
  flex: 1;
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 2.5rem;
}

.release-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
}

.release-body {
  color: var(--text-muted);
  line-height: 1.7;
  font-size: 1.05rem;
}

.fade-in-up {
  opacity: 0;
  transform: translateY(20px);
  animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.delay-1 { animation-delay: 0.1s; }
.delay-2 { animation-delay: 0.2s; }

@keyframes fadeInUp {
  to { opacity: 1; transform: translateY(0); }
}

@media (max-width: 768px) {
  .page-container { padding: 4rem 2rem; }
  .page-title { font-size: 2.5rem; }
  
  .changelog-list::before {
    left: 20px;
  }
  
  .release-item {
    flex-direction: column;
    gap: 1.5rem;
  }
  
  .release-meta {
    width: 100%;
    text-align: left;
    display: flex;
    align-items: center;
    gap: 1rem;
    padding-left: 45px;
  }
  
  .release-meta::after {
    left: 15px;
    right: auto;
  }
}
</style>
