<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'

useSeo({
  title: 'Integrations - CloudySMS',
  description: 'Connect CloudySMS with your favorite tools. Integrate with CRMs, payment gateways, and backend systems easily.'
})

const integrations = [
  { name: 'Salesforce', desc: 'Sync leads and WhatsApp conversations directly to Salesforce CRM.', color: '#00a1e0' },
  { name: 'HubSpot', desc: 'Automate marketing workflows based on WhatsApp engagement.', color: '#ff7a59' },
  { name: 'Shopify', desc: 'Send order updates and abandoned cart recovery messages.', color: '#95bf47' },
  { name: 'Stripe', desc: 'Process payments and send payment links directly in WhatsApp.', color: '#635bff' },
  { name: 'Zapier', desc: 'Connect to 5,000+ apps without writing a single line of code.', color: '#ff4f00' },
  { name: 'Custom Webhooks', desc: 'Trigger external APIs and receive real-time event payloads.', color: '#10b981' }
]
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Connect your <span class="text-gradient">stack</span></h1>
        <p class="page-subtitle fade-in-up delay-1">Integrate CloudySMS seamlessly with the tools you already use every day.</p>
      </div>

      <div class="integrations-grid fade-in-up delay-2">
        <div class="integration-card" v-for="item in integrations" :key="item.name">
          <div class="integration-logo" :style="{ backgroundColor: item.color + '22', color: item.color }">
            {{ item.name.charAt(0) }}
          </div>
          <h3 class="integration-name">{{ item.name }}</h3>
          <p class="integration-desc">{{ item.desc }}</p>
          <a href="#" class="integration-link">Learn more &rarr;</a>
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

.text-gradient {
  color: var(--text-main);
}

.page-subtitle {
  font-size: 1.25rem;
  color: var(--text-muted);
  max-width: 600px;
  margin: 0 auto;
}

.integrations-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
}

.integration-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 2.5rem;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
}

.integration-card:hover {
  transform: translateY(-5px);
  background: var(--bg-color-alt);
  border-color: #3f3f46;
}

.integration-logo {
  width: 4rem;
  height: 4rem;
  border-radius: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
}

.integration-name {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
}

.integration-desc {
  color: var(--text-muted);
  line-height: 1.6;
  font-size: 1rem;
  margin-bottom: 2rem;
  flex: 1;
}

.integration-link {
  color: #10b981;
  text-decoration: none;
  font-weight: 500;
  margin-top: auto;
  transition: color 0.2s;
}

.integration-link:hover {
  color: #34d399;
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
  .integrations-grid { grid-template-columns: 1fr; }
}
</style>
