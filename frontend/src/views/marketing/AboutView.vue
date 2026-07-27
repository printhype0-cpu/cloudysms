<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'

useSeo({
  title: 'About - CloudySMS',
  description: 'Learn about our mission to democratize conversational commerce with the open-source WhatsApp Business Platform.'
})

const team = [
  { name: 'Alex Rivera', role: 'CEO & Founder', image: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
  { name: 'Sarah Chen', role: 'CTO', image: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
  { name: 'Marcus Johnson', role: 'Head of Product', image: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
  { name: 'Elena Rodriguez', role: 'Lead Designer', image: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' }
]
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Our <span class="text-gradient">Mission</span></h1>
        <p class="page-subtitle fade-in-up delay-1">We are building the open-source standard for conversational commerce and customer support on WhatsApp.</p>
      </div>

      <div class="about-content fade-in-up delay-2">
        <div class="story-section">
          <h2>The Story</h2>
          <p>We started CloudySMS because we were frustrated by the complexity and high costs of existing WhatsApp Business solutions. We believed that businesses of all sizes should have access to enterprise-grade communication tools without being locked into restrictive, expensive platforms.</p>
          <p>Today, CloudySMS powers millions of conversations worldwide, helping companies build deeper relationships with their customers through automation, AI, and seamless integrations.</p>
        </div>

        <div class="team-section">
          <h2>Meet the Team</h2>
          <div class="team-grid">
            <div class="team-member" v-for="member in team" :key="member.name">
              <img :src="member.image" :alt="member.name" class="member-image" />
              <h3 class="member-name">{{ member.name }}</h3>
              <p class="member-role">{{ member.role }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </MarketingLayout>
</template>

<style scoped>
.page-container {
  padding: 6rem 4rem;
  max-width: 1000px;
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

.about-content h2 {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
}

.story-section {
  margin-bottom: 5rem;
}

.story-section p {
  color: var(--text-muted);
  font-size: 1.1rem;
  line-height: 1.8;
  margin-bottom: 1.5rem;
}

.team-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 2rem;
  margin-top: 2rem;
}

.team-member {
  text-align: center;
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 2rem;
  transition: transform 0.3s ease;
}

.team-member:hover {
  transform: translateY(-5px);
  background: var(--bg-color-alt);
}

.member-image {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  margin-bottom: 1.5rem;
  border: 2px solid rgba(16, 185, 129, 0.3);
}

.member-name {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.member-role {
  color: #10b981;
  font-size: 0.95rem;
  font-weight: 500;
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
