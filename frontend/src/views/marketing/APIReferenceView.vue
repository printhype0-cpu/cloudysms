<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'
import { Copy, Check } from 'lucide-vue-next'
import { ref } from 'vue'

useSeo({
  title: 'API Reference - CloudySMS',
  description: 'Explore the CloudySMS REST API. Send messages, manage contacts, and build custom integrations.'
})

const copied = ref(false)
const codeSnippet = `curl -X POST https://api.cloudysms.com/v1/messages \\
  -H "Authorization: Bearer YOUR_API_KEY" \\
  -H "Content-Type: application/json" \\
  -d '{
    "to": "+1234567890",
    "type": "text",
    "text": {
      "body": "Hello from CloudySMS API!"
    }
  }'`

const copyCode = () => {
  navigator.clipboard.writeText(codeSnippet)
  copied.value = true
  setTimeout(() => { copied.value = false }, 2000)
}
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="api-hero fade-in-up">
        <div class="api-hero-content">
          <h1 class="page-title">API <span class="text-gradient">Reference</span></h1>
          <p class="page-subtitle">Build powerful WhatsApp integrations with our REST API. Send messages, manage webhooks, and automate workflows programmatically.</p>
          <div class="hero-actions">
            <button class="btn btn-primary">Generate API Key</button>
            <button class="btn btn-outline">Read Full Docs</button>
          </div>
        </div>
        
        <div class="code-window fade-in-up delay-1">
          <div class="window-header">
            <div class="dots">
              <span class="dot red"></span>
              <span class="dot yellow"></span>
              <span class="dot green"></span>
            </div>
            <div class="window-title">Send Message Request</div>
            <button class="copy-btn" @click="copyCode" :class="{ 'copied': copied }">
              <Check v-if="copied" class="icon-small" />
              <Copy v-else class="icon-small" />
            </button>
          </div>
          <pre><code>{{ codeSnippet }}</code></pre>
        </div>
      </div>

      <div class="endpoints-section fade-in-up delay-2">
        <h2>Popular Endpoints</h2>
        <div class="endpoints-grid">
          <div class="endpoint-card">
            <div class="method post">POST</div>
            <div class="path">/v1/messages</div>
            <p class="endpoint-desc">Send text, media, or template messages to a contact.</p>
          </div>
          <div class="endpoint-card">
            <div class="method get">GET</div>
            <div class="path">/v1/contacts</div>
            <p class="endpoint-desc">Retrieve a paginated list of your contacts.</p>
          </div>
          <div class="endpoint-card">
            <div class="method post">POST</div>
            <div class="path">/v1/webhooks</div>
            <p class="endpoint-desc">Register a new webhook URL for incoming messages.</p>
          </div>
          <div class="endpoint-card">
            <div class="method get">GET</div>
            <div class="path">/v1/analytics/overview</div>
            <p class="endpoint-desc">Fetch aggregate metrics for messages sent and delivered.</p>
          </div>
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

.api-hero {
  display: flex;
  align-items: center;
  gap: 4rem;
  margin-bottom: 6rem;
}

.api-hero-content {
  flex: 1;
  max-width: 600px;
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
  line-height: 1.6;
  margin-bottom: 2.5rem;
}

.hero-actions {
  display: flex;
  gap: 1rem;
}

.btn {
  padding: 0.875rem 2rem;
  border-radius: 0.6rem;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
}

.btn-primary {
  background: linear-gradient(135deg, #10b981, #059669);
  color: var(--text-main);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(16, 185, 129, 0.35);
}

.btn-outline {
  background: transparent;
  color: var(--text-main);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.btn-outline:hover {
  background: var(--bg-color-alt);
}

.code-window {
  flex: 1;
  background: #0d0d0f;
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
}

.window-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 1.5rem;
  background: var(--glass-bg);
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}

.dots {
  display: flex;
  gap: 0.5rem;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}
.dot.red { background: #ef4444; }
.dot.yellow { background: #eab308; }
.dot.green { background: #22c55e; }

.window-title {
  color: var(--text-muted);
  font-size: 0.9rem;
  font-family: monospace;
}

.copy-btn {
  background: transparent;
  border: none;
  color: var(--text-muted);
  cursor: pointer;
  transition: color 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.copy-btn:hover {
  color: var(--text-main);
}

.copy-btn.copied {
  color: #10b981;
}

.icon-small {
  width: 1.25rem;
  height: 1.25rem;
}

pre {
  margin: 0;
  padding: 1.5rem;
  overflow-x: auto;
}

code {
  font-family: 'Fira Code', 'Courier New', Courier, monospace;
  font-size: 0.95rem;
  color: var(--text-main);
  line-height: 1.6;
}

.endpoints-section h2 {
  font-size: 2rem;
  margin-bottom: 2rem;
}

.endpoints-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 1.5rem;
}

.endpoint-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 0.75rem;
  padding: 1.5rem;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  transition: transform 0.2s;
}

.endpoint-card:hover {
  transform: translateX(5px);
  background: var(--bg-color-alt);
}

.method {
  font-size: 0.85rem;
  font-weight: 700;
  padding: 0.25rem 0.75rem;
  border-radius: 0.5rem;
  letter-spacing: 0.05em;
}

.method.post { background: rgba(16, 185, 129, 0.15); color: #10b981; }
.method.get { background: rgba(59, 130, 246, 0.15); color: #3b82f6; }

.path {
  font-family: monospace;
  font-size: 1.1rem;
  color: var(--text-main);
  min-width: 150px;
}

.endpoint-desc {
  color: var(--text-muted);
  font-size: 0.95rem;
  margin: 0;
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

@media (max-width: 1024px) {
  .api-hero { flex-direction: column; }
  .code-window { width: 100%; }
}

@media (max-width: 768px) {
  .page-container { padding: 4rem 2rem; }
  .page-title { font-size: 2.5rem; }
  .endpoint-card { flex-direction: column; align-items: flex-start; gap: 1rem; }
  .path { min-width: auto; }
}
</style>
