<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'
import { MessageCircle, Github, Calendar, Users } from 'lucide-vue-next'

useSeo({
  title: 'Community - CloudySMS',
  description: 'Join the open-source CloudySMS community. Contribute, connect, and collaborate with developers worldwide.'
})
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Join the <span class="text-gradient">Community</span></h1>
        <p class="page-subtitle fade-in-up delay-1">Connect with developers, get help, and contribute to the open-source WhatsApp platform.</p>
      </div>

      <div class="community-grid fade-in-up delay-2">
        <div class="community-card discord">
          <div class="card-icon">
            <MessageCircle class="icon" />
          </div>
          <h3 class="card-title">Discord Server</h3>
          <p class="card-desc">Chat with the core team and over 5,000 developers. Get real-time help, share your projects, and stay updated.</p>
          <button class="btn btn-primary w-full">Join Discord</button>
        </div>
        
        <div class="community-card github">
          <div class="card-icon">
            <Github class="icon" />
          </div>
          <h3 class="card-title">GitHub Repository</h3>
          <p class="card-desc">CloudySMS is open-source. Browse the code, submit pull requests, report issues, and star our repository.</p>
          <button class="btn btn-outline w-full">View on GitHub</button>
        </div>
        
        <div class="community-card events">
          <div class="card-icon">
            <Calendar class="icon" />
          </div>
          <h3 class="card-title">Upcoming Events</h3>
          <p class="card-desc">Join our monthly town halls, technical workshops, and Q&A sessions. See what we are building next.</p>
          <button class="btn btn-outline w-full">View Schedule</button>
        </div>
        
        <div class="community-card forum">
          <div class="card-icon">
            <Users class="icon" />
          </div>
          <h3 class="card-title">Discussion Forums</h3>
          <p class="card-desc">Dive deep into architectural discussions, propose new features, and browse community tutorials.</p>
          <button class="btn btn-outline w-full">Visit Forums</button>
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

.community-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.community-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1.25rem;
  padding: 2.5rem;
  display: flex;
  flex-direction: column;
  transition: transform 0.3s ease;
}

.community-card:hover {
  transform: translateY(-5px);
}

.discord { border-top: 4px solid #5865F2; }
.github { border-top: 4px solid #f0f6fc; }
.events { border-top: 4px solid #10b981; }
.forum { border-top: 4px solid #3b82f6; }

.card-icon {
  width: 4rem;
  height: 4rem;
  border-radius: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-color-alt);
  margin-bottom: 1.5rem;
}

.icon {
  width: 2rem;
  height: 2rem;
  color: var(--text-main);
}

.card-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1rem;
}

.card-desc {
  color: var(--text-muted);
  line-height: 1.6;
  margin-bottom: 2rem;
  flex: 1;
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
  background: #5865F2; /* Discord color */
  color: var(--text-main);
}

.btn-primary:hover {
  background: #4752C4;
}

.btn-outline {
  background: transparent;
  color: var(--text-main);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.btn-outline:hover {
  background: var(--bg-color-alt);
}

.w-full { width: 100%; }

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
