<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'
import { Bot, BarChart3, Users, Zap, Shield, Globe, MessageSquare, Workflow } from 'lucide-vue-next'

useSeo({
  title: 'Features - CloudySMS',
  description: 'Explore the powerful features of CloudySMS. Build chatbots, manage team inboxes, and scale your WhatsApp communication.'
})
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Powerful <span class="text-gradient">Features</span></h1>
        <p class="page-subtitle fade-in-up delay-1">Everything you need to build, manage, and scale your WhatsApp operations.</p>
      </div>

      <div class="features-grid fade-in-up delay-2">
        <div class="feature-card">
          <div class="feature-icon-wrapper blue">
            <Bot class="feature-icon" />
          </div>
          <h3 class="feature-title">Visual Flow Builder</h3>
          <p class="feature-desc">Design complex conversational flows with a drag-and-drop interface. No coding required.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper green">
            <Users class="feature-icon" />
          </div>
          <h3 class="feature-title">Team Inbox</h3>
          <p class="feature-desc">Collaborate on customer queries in real-time. Assign tickets, leave internal notes, and resolve issues faster.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper purple">
            <BarChart3 class="feature-icon" />
          </div>
          <h3 class="feature-title">Rich Analytics</h3>
          <p class="feature-desc">Track agent performance, message delivery rates, and customer satisfaction with detailed insights.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper orange">
            <Zap class="feature-icon" />
          </div>
          <h3 class="feature-title">Automated Campaigns</h3>
          <p class="feature-desc">Send personalized broadcast messages to segments of your audience and track engagement.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper red">
            <Shield class="feature-icon" />
          </div>
          <h3 class="feature-title">Enterprise Security</h3>
          <p class="feature-desc">Role-based access control, SSO integration, and detailed audit logs to keep your data secure.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper cyan">
            <Globe class="feature-icon" />
          </div>
          <h3 class="feature-title">API & Webhooks</h3>
          <p class="feature-desc">Integrate seamlessly with your existing CRM or backend systems using our robust developer API.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper yellow">
            <MessageSquare class="feature-icon" />
          </div>
          <h3 class="feature-title">Template Management</h3>
          <p class="feature-desc">Create, submit, and manage WhatsApp approved templates directly from the dashboard.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper pink">
            <Workflow class="feature-icon" />
          </div>
          <h3 class="feature-title">Custom Actions</h3>
          <p class="feature-desc">Trigger external API calls and background jobs directly from within a chatbot flow.</p>
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

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.feature-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 2.5rem;
  transition: all 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-5px);
  background: var(--bg-color-alt);
  border-color: #3f3f46;
}

.feature-icon-wrapper {
  width: 3.5rem;
  height: 3.5rem;
  border-radius: 0.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 1.5rem;
}

.feature-icon-wrapper.blue { background: rgba(59, 130, 246, 0.1); color: #3b82f6; }
.feature-icon-wrapper.green { background: rgba(16, 185, 129, 0.1); color: #10b981; }
.feature-icon-wrapper.purple { background: rgba(139, 92, 246, 0.1); color: #8b5cf6; }
.feature-icon-wrapper.orange { background: rgba(249, 115, 22, 0.1); color: #f97316; }
.feature-icon-wrapper.red { background: rgba(239, 68, 68, 0.1); color: #ef4444; }
.feature-icon-wrapper.cyan { background: rgba(6, 182, 212, 0.1); color: #06b6d4; }
.feature-icon-wrapper.yellow { background: rgba(234, 179, 8, 0.1); color: #eab308; }
.feature-icon-wrapper.pink { background: rgba(236, 72, 153, 0.1); color: #ec4899; }

.feature-icon {
  width: 1.75rem;
  height: 1.75rem;
}

.feature-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
}

.feature-desc {
  color: var(--text-muted);
  line-height: 1.6;
  font-size: 0.95rem;
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
}
</style>
