<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'
import { BookOpen, FileCode2, PlayCircle, Settings, Key, Shield } from 'lucide-vue-next'

useSeo({
  title: 'Documentation - CloudySMS',
  description: 'Learn how to integrate, configure, and scale with CloudySMS through our comprehensive guides and tutorials.'
})

const sections = [
  {
    title: 'Getting Started',
    icon: PlayCircle,
    color: '#10b981',
    links: ['Platform Overview', 'Quickstart Guide', 'Connecting WhatsApp', 'Inviting Team Members']
  },
  {
    title: 'Core Concepts',
    icon: BookOpen,
    color: '#3b82f6',
    links: ['Understanding Flows', 'Managing Contacts', 'Tags and Segments', 'Templates Overview']
  },
  {
    title: 'Development',
    icon: FileCode2,
    color: '#8b5cf6',
    links: ['REST API Basics', 'Webhooks Integration', 'SDK Installation', 'Custom Actions']
  },
  {
    title: 'Account Settings',
    icon: Settings,
    color: '#f97316',
    links: ['Billing and Plans', 'Custom Domains', 'Notification Preferences', 'Data Export']
  },
  {
    title: 'Security',
    icon: Shield,
    color: '#ef4444',
    links: ['Role-Based Access', 'Audit Logs', 'Compliance', 'Data Retention']
  },
  {
    title: 'Authentication',
    icon: Key,
    color: '#06b6d4',
    links: ['API Keys', 'SSO Setup', 'OAuth 2.0', 'Token Management']
  }
]
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Documentation</h1>
        <p class="page-subtitle fade-in-up delay-1">Everything you need to build and scale with CloudySMS.</p>
        
        <div class="search-bar fade-in-up delay-2">
          <input type="text" placeholder="Search documentation..." class="search-input" />
          <button class="btn btn-primary search-btn">Search</button>
        </div>
      </div>

      <div class="docs-grid fade-in-up delay-3">
        <div class="docs-card" v-for="section in sections" :key="section.title">
          <div class="card-header">
            <div class="icon-wrapper" :style="{ backgroundColor: section.color + '22', color: section.color }">
              <component :is="section.icon" class="icon" />
            </div>
            <h3 class="card-title">{{ section.title }}</h3>
          </div>
          <ul class="docs-links">
            <li v-for="link in section.links" :key="link">
              <a href="#">{{ link }}</a>
            </li>
          </ul>
          <a href="#" class="view-all">View all &rarr;</a>
        </div>
      </div>
    </div>
  </MarketingLayout>
</template>

<style scoped>
.page-container {
  padding: 6rem 4rem;
  max-width: 1400px;
  margin: 0 auto;
  width: 100%;
}

.page-header {
  text-align: center;
  margin-bottom: 5rem;
}

.page-title {
  font-size: 3.5rem;
  font-weight: 800;
  margin-bottom: 1.5rem;
  letter-spacing: -0.03em;
}

.page-subtitle {
  font-size: 1.25rem;
  color: var(--text-muted);
  max-width: 600px;
  margin: 0 auto 2.5rem;
}

.search-bar {
  display: flex;
  max-width: 600px;
  margin: 0 auto;
  gap: 1rem;
}

.search-input {
  flex: 1;
  padding: 1rem 1.5rem;
  border-radius: 0.75rem;
  background: var(--bg-color-alt);
  border: 1px solid var(--glass-border);
  color: var(--text-main);
  font-size: 1rem;
  outline: none;
  transition: border-color 0.2s;
}

.search-input:focus {
  border-color: #10b981;
}

.search-btn {
  padding: 0 2rem;
  border-radius: 0.75rem;
}

.btn-primary {
  background: linear-gradient(135deg, #10b981, #059669);
  color: var(--text-main);
  border: none;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s;
}

.btn-primary:hover {
  transform: translateY(-2px);
}

.docs-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
}

.docs-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 2.5rem;
  display: flex;
  flex-direction: column;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.icon-wrapper {
  width: 3rem;
  height: 3rem;
  border-radius: 0.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon {
  width: 1.5rem;
  height: 1.5rem;
}

.card-title {
  font-size: 1.25rem;
  font-weight: 600;
}

.docs-links {
  list-style: none;
  padding: 0;
  margin: 0 0 1.5rem 0;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  flex: 1;
}

.docs-links a {
  color: var(--text-muted);
  text-decoration: none;
  transition: color 0.2s;
}

.docs-links a:hover {
  color: #10b981;
}

.view-all {
  color: #10b981;
  text-decoration: none;
  font-weight: 500;
  font-size: 0.95rem;
  transition: color 0.2s;
}

.view-all:hover {
  color: #34d399;
}

.fade-in-up {
  opacity: 0;
  transform: translateY(20px);
  animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.delay-1 { animation-delay: 0.1s; }
.delay-2 { animation-delay: 0.2s; }
.delay-3 { animation-delay: 0.3s; }

@keyframes fadeInUp {
  to { opacity: 1; transform: translateY(0); }
}

@media (max-width: 768px) {
  .page-container { padding: 4rem 2rem; }
  .page-title { font-size: 2.5rem; }
  .search-bar { flex-direction: column; }
}
</style>
