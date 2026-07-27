<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'

useSeo({
  title: 'Blog - CloudySMS',
  description: 'Read the latest news, tutorials, and insights from the CloudySMS team.'
})

const posts = [
  {
    title: 'Building a Customer Support Chatbot in 10 Minutes',
    category: 'Tutorial',
    date: 'August 12, 2026',
    author: 'Sarah Jenkins',
    excerpt: 'Learn how to use the new Visual Flow Builder to automate your level 1 customer support and deflect tickets efficiently.',
    image: 'https://images.unsplash.com/photo-1551288049-bebda4e38f71?auto=format&fit=crop&w=800&q=80',
    color: '#3b82f6'
  },
  {
    title: 'The Future of Conversational Commerce',
    category: 'Industry Insight',
    date: 'August 02, 2026',
    author: 'David Chen',
    excerpt: 'Why WhatsApp is becoming the ultimate storefront for modern businesses, and how you can prepare for the conversational revolution.',
    image: 'https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?auto=format&fit=crop&w=800&q=80',
    color: '#8b5cf6'
  },
  {
    title: 'Best Practices for WhatsApp Broadcast Campaigns',
    category: 'Marketing',
    date: 'July 20, 2026',
    author: 'Amanda Clark',
    excerpt: 'Avoid getting blocked by users. Discover the golden rules of sending marketing messages that people actually want to read.',
    image: 'https://images.unsplash.com/photo-1432821596592-e2c18b78144f?auto=format&fit=crop&w=800&q=80',
    color: '#10b981'
  }
]
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">The <span class="text-gradient">Blog</span></h1>
        <p class="page-subtitle fade-in-up delay-1">Insights, tutorials, and updates from the CloudySMS team.</p>
      </div>

      <div class="blog-grid fade-in-up delay-2">
        <div class="blog-card" v-for="post in posts" :key="post.title">
          <div class="blog-image" :style="{ backgroundImage: `url(${post.image})` }">
            <div class="category-badge" :style="{ backgroundColor: post.color }">{{ post.category }}</div>
          </div>
          <div class="blog-content">
            <div class="blog-meta">
              <span>{{ post.date }}</span>
              <span class="dot">•</span>
              <span>{{ post.author }}</span>
            </div>
            <h3 class="blog-title"><a href="#">{{ post.title }}</a></h3>
            <p class="blog-excerpt">{{ post.excerpt }}</p>
            <a href="#" class="read-more">Read article &rarr;</a>
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

.blog-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2.5rem;
}

.blog-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1.25rem;
  overflow: hidden;
  transition: transform 0.3s ease, border-color 0.3s ease;
  display: flex;
  flex-direction: column;
}

.blog-card:hover {
  transform: translateY(-5px);
  border-color: #3f3f46;
}

.blog-image {
  height: 200px;
  background-size: cover;
  background-position: center;
  position: relative;
}

.category-badge {
  position: absolute;
  top: 1rem;
  right: 1rem;
  color: var(--text-main);
  padding: 0.25rem 0.75rem;
  border-radius: 2rem;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.blog-content {
  padding: 2rem;
  display: flex;
  flex-direction: column;
  flex: 1;
}

.blog-meta {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--text-muted);
  font-size: 0.85rem;
  margin-bottom: 1rem;
}

.dot {
  font-size: 1.2rem;
  line-height: 1;
}

.blog-title {
  font-size: 1.4rem;
  font-weight: 700;
  margin-bottom: 1rem;
  line-height: 1.4;
}

.blog-title a {
  color: var(--text-main);
  text-decoration: none;
  transition: color 0.2s;
}

.blog-title a:hover {
  color: #10b981;
}

.blog-excerpt {
  color: var(--text-muted);
  line-height: 1.6;
  font-size: 0.95rem;
  margin-bottom: 2rem;
  flex: 1;
}

.read-more {
  color: #10b981;
  text-decoration: none;
  font-weight: 500;
  margin-top: auto;
  transition: color 0.2s;
}

.read-more:hover {
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
}
</style>
