<script setup lang="ts">
import { useRouter } from 'vue-router'
import { Bot, BarChart3, Users, Zap, CheckCircle2, Globe, Shield, Smartphone } from 'lucide-vue-next'
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'
import HeroSection from '@/components/HeroSection.vue'

const router = useRouter()

// Add SEO tags
useSeo({
  title: 'Home',
  description: 'CloudySMS - The Modern, Open-Source WhatsApp Business Platform. Automate conversations, build chatbots, and engage your customers at scale.',
  image: 'https://www.cloudysms.com/og-image.png',
  schemas: [
    {
      "@context": "https://schema.org",
      "@type": "BreadcrumbList",
      "itemListElement": [
        {
          "@type": "ListItem",
          "position": 1,
          "name": "Home",
          "item": "https://www.cloudysms.com/"
        }
      ]
    },
    {
      "@context": "https://schema.org",
      "@type": "Organization",
      "name": "CloudySMS",
      "url": "https://www.cloudysms.com",
      "logo": "https://www.cloudysms.com/favicon.svg"
    },
    {
      "@context": "https://schema.org",
      "@type": "WebPage",
      "name": "Turn WhatsApp into Your #1 Sales Channel | CloudySMS",
      "description": "CloudySMS - The Modern, Open-Source WhatsApp Business Platform. Automate conversations, build chatbots, and engage your customers at scale.",
      "publisher": {
        "@type": "Organization",
        "name": "CloudySMS"
      }
    },
    {
      "@context": "https://schema.org",
      "@type": "FAQPage",
      "mainEntity": [
        {
          "@type": "Question",
          "name": "What is CloudySMS?",
          "acceptedAnswer": {
            "@type": "Answer",
            "text": "CloudySMS is a modern, open-source WhatsApp Business Platform designed for marketing and customer support teams to automate conversations, build chatbots, and send broadcast campaigns at scale."
          }
        },
        {
          "@type": "Question",
          "name": "How to automate WhatsApp messages?",
          "acceptedAnswer": {
            "@type": "Answer",
            "text": "With CloudySMS, you can automate WhatsApp messages using our Visual Flow Builder. Simply drag and drop conversation blocks to create keyword-based auto-replies or AI-powered chatbots without writing any code."
          }
        },
        {
          "@type": "Question",
          "name": "Who is CloudySMS best for?",
          "acceptedAnswer": {
            "@type": "Answer",
            "text": "CloudySMS is best for e-commerce businesses, sales teams, and customer support centers that need to handle high volumes of WhatsApp messages, manage a shared team inbox, or run targeted marketing campaigns."
          }
        }
      ]
    }
  ]
})

const getStarted = () => {
  router.push('/login')
}
</script>

<template>
  <MarketingLayout>
    <!-- Hero Section -->
    <HeroSection />

    <!-- Top Summary Answer (AEO Signal) -->
    <section class="section-container pt-0 pb-12 max-w-4xl mx-auto">
      <div class="p-6 md:p-8 bg-slate-50/50 rounded-2xl border border-slate-100/50 backdrop-blur-sm text-left">
        <h2 class="text-xl md:text-2xl font-bold text-[color:var(--text-main)] mb-3">Key Takeaway</h2>
        <ul class="list-disc pl-5 text-base md:text-lg text-[color:var(--text-muted)] leading-relaxed space-y-2">
          <li><strong>CloudySMS</strong> is the complete WhatsApp Business Platform designed for growing sales and support teams.</li>
          <li>According to <a href="https://business.whatsapp.com/" target="_blank" rel="noopener noreferrer" class="text-emerald-600 hover:underline">official WhatsApp Business</a> reports, messages achieve up to a 98% open rate, making it an essential sales channel.</li>
          <li>By combining a shared team inbox, visual chatbot builder, and automated marketing campaigns, it helps you convert leads faster and deliver 24/7 customer support.</li>
          <li>Manage everything from one unified <router-link to="/dashboard" class="text-emerald-600 hover:underline">dashboard</router-link>.</li>
        </ul>
      </div>
    </section>

    <!-- Features Section -->
    <section id="features" class="section-container">
      <div class="section-header text-center">
        <h2 class="section-title">Everything you need to <span class="text-gradient">engage</span></h2>
        <p class="section-subtitle">Powerful tools designed for modern customer support and marketing teams.</p>
      </div>
      <div class="features-grid">
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
          <div class="feature-icon-wrapper purple">
            <Smartphone class="feature-icon" />
          </div>
          <h3 class="feature-title">Mobile App</h3>
          <p class="feature-desc">Manage customer conversations on the go with our dedicated iOS and Android applications.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon-wrapper green">
            <CheckCircle2 class="feature-icon" />
          </div>
          <h3 class="feature-title">Guaranteed Uptime</h3>
          <p class="feature-desc">Enjoy a 99.99% uptime SLA to ensure your business never misses an important customer message.</p>
        </div>
      </div>
    </section>

    <!-- Solutions Section -->
    <section id="solutions" class="section-container alternate-bg">
      <div class="solutions-content">
        <div class="solutions-text">
          <h2 class="section-title">Built for <span class="text-gradient">Scale</span></h2>
          <p class="section-subtitle text-left">Whether you're a growing startup or an established enterprise, CloudySMS adapts to your communication needs.</p>
          <ul class="solutions-list">
            <li>
              <CheckCircle2 class="list-icon" />
              <div>
                <strong>Marketing & Sales Teams:</strong> 
                <span class="block text-sm mt-1">Drive conversions with targeted WhatsApp broadcast campaigns, recover abandoned carts, and qualify leads automatically.</span>
              </div>
            </li>
            <li>
              <CheckCircle2 class="list-icon" />
              <div>
                <strong>Customer Support:</strong> 
                <span class="block text-sm mt-1">Provide instant 24/7 support using AI chatbots and manage complex queries with our collaborative team inbox.</span>
              </div>
            </li>
            <li>
              <CheckCircle2 class="list-icon" />
              <div>
                <strong>E-commerce & Operations:</strong> 
                <span class="block text-sm mt-1">Send automated transactional alerts, order updates, and appointment reminders using official <a href="https://developers.facebook.com/docs/whatsapp/" target="_blank" rel="noopener noreferrer" class="text-emerald-600 hover:underline">Meta API</a> templates.</span>
              </div>
            </li>
          </ul>
        </div>
        <div class="solutions-visual">
          <div class="glass-panel showcase-panel">
             <div class="showcase-item">
                <Smartphone class="showcase-icon" />
                <div class="showcase-lines">
                  <div class="s-line w-full"></div>
                  <div class="s-line w-3/4"></div>
                </div>
             </div>
             <div class="showcase-item right">
                <div class="showcase-lines right-align">
                  <div class="s-line w-full"></div>
                  <div class="s-line w-1/2"></div>
                </div>
                <div class="user-avatar"></div>
             </div>
             <div class="showcase-item">
                <Bot class="showcase-icon" />
                <div class="showcase-lines">
                  <div class="s-line w-full"></div>
                  <div class="s-line w-5/6"></div>
                </div>
             </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Pricing Section -->
    <section id="pricing" class="section-container">
      <div class="section-header text-center">
        <h2 class="section-title">Simple, transparent <span class="text-gradient">pricing</span></h2>
        <p class="section-subtitle">Start for free, upgrade when you need more power.</p>
      </div>
      <div class="pricing-grid">
        <!-- Starter Plan -->
        <div class="pricing-card">
          <h3 class="plan-name">Starter</h3>
          <div class="plan-price">
            <span class="currency">$</span><span class="amount">0</span><span class="period">/mo</span>
          </div>
          <p class="plan-desc">Perfect for small teams getting started.</p>
          <ul class="plan-features">
            <li><CheckCircle2 class="check" /> Up to 1,000 messages/mo</li>
            <li><CheckCircle2 class="check" /> 2 Team Members</li>
            <li><CheckCircle2 class="check" /> Basic Chatbot Builder</li>
            <li><CheckCircle2 class="check" /> Community Support</li>
          </ul>
          <button class="btn btn-outline w-full mt-auto" @click="getStarted">Get Started</button>
        </div>
        <!-- Pro Plan -->
        <div class="pricing-card popular">
          <div class="popular-badge">Most Popular</div>
          <h3 class="plan-name">Pro</h3>
          <div class="plan-price">
            <span class="currency">$</span><span class="amount">49</span><span class="period">/mo</span>
          </div>
          <p class="plan-desc">Everything you need for growing businesses.</p>
          <ul class="plan-features">
            <li><CheckCircle2 class="check" /> 10,000 messages/mo</li>
            <li><CheckCircle2 class="check" /> Unlimited Team Members</li>
            <li><CheckCircle2 class="check" /> Advanced AI Chatbots</li>
            <li><CheckCircle2 class="check" /> API & Webhooks</li>
            <li><CheckCircle2 class="check" /> Priority Support</li>
          </ul>
          <button class="btn btn-primary w-full mt-auto" @click="getStarted">Start Free Trial</button>
        </div>
        <!-- Enterprise Plan -->
        <div class="pricing-card">
          <h3 class="plan-name">Enterprise</h3>
          <div class="plan-price">
            <span class="amount">Custom</span>
          </div>
          <p class="plan-desc">For large scale organizations and custom needs.</p>
          <ul class="plan-features">
            <li><CheckCircle2 class="check" /> Unlimited messages</li>
            <li><CheckCircle2 class="check" /> Custom Integrations</li>
            <li><CheckCircle2 class="check" /> Dedicated Account Manager</li>
            <li><CheckCircle2 class="check" /> SLA Guarantee</li>
            <li><CheckCircle2 class="check" /> Self-hosted Option</li>
          </ul>
          <button class="btn btn-outline w-full mt-auto" @click="getStarted">Contact Sales</button>
        </div>
      </div>
    </section>

    <!-- FAQ Section -->
    <section id="faq" class="section-container alternate-bg">
      <div class="section-header text-center">
        <h2 class="section-title">Frequently Asked <span class="text-gradient">Questions</span></h2>
        <p class="section-subtitle">Everything you need to know about getting started with CloudySMS.</p>
      </div>
      <div class="max-w-3xl mx-auto flex flex-col gap-6">
        
        <div class="bg-white p-6 rounded-xl border border-[color:var(--glass-border)] shadow-sm">
          <h3 class="text-xl font-bold mb-3 text-[color:var(--text-main)]">What is CloudySMS?</h3>
          <p class="text-[color:var(--text-muted)] leading-relaxed">
            <strong>Short answer:</strong> CloudySMS is a modern, open-source WhatsApp Business Platform designed for marketing and customer support teams to automate conversations, build chatbots, and send broadcast campaigns at scale.
          </p>
        </div>

        <div class="bg-white p-6 rounded-xl border border-[color:var(--glass-border)] shadow-sm">
          <h3 class="text-xl font-bold mb-3 text-[color:var(--text-main)]">How to automate WhatsApp messages?</h3>
          <p class="text-[color:var(--text-muted)] leading-relaxed mb-3">
            <strong>Short answer:</strong> With CloudySMS, you can automate WhatsApp messages using our Visual Flow Builder without writing any code. Follow these steps:
          </p>
          <ul class="list-decimal pl-5 text-[color:var(--text-muted)] space-y-2">
            <li>Connect your WhatsApp Business API account.</li>
            <li>Navigate to the <router-link to="/chatbot" class="text-emerald-600 hover:underline">Automations</router-link> tab in your dashboard.</li>
            <li>Use the drag-and-drop builder to create keyword-triggered auto-replies.</li>
            <li>Publish your flow to start engaging customers instantly.</li>
          </ul>
        </div>

        <div class="bg-white p-6 rounded-xl border border-[color:var(--glass-border)] shadow-sm">
          <h3 class="text-xl font-bold mb-3 text-[color:var(--text-main)]">Who is CloudySMS best for?</h3>
          <p class="text-[color:var(--text-muted)] leading-relaxed">
            <strong>Short answer:</strong> CloudySMS is best for e-commerce businesses, sales teams, and customer support centers that need to handle high volumes of WhatsApp messages, manage a shared team inbox, or run targeted marketing campaigns. It offers robust tools compared to the standard WhatsApp Business app.
          </p>
        </div>

      </div>
    </section>

    <!-- Ownership & Footer -->
    <footer class="section-container pt-12 pb-8 border-t border-[color:var(--glass-border)]">
      <div class="max-w-6xl mx-auto flex flex-col md:flex-row justify-between items-center gap-6">
        <div>
          <div class="logo mb-2">
            <div class="logo-icon small">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M3 21l1.65-3.8a9 9 0 1 1 3.4 2.9L3 21"></path>
                <path d="M9 10a.5.5 0 0 0 1 0V9a.5.5 0 0 0-1 0v1a5 5 0 0 0 5 5h1a.5.5 0 0 0 0-1h-1a.5.5 0 0 0 0 1"></path>
              </svg>
            </div>
            CloudySMS
          </div>
          <p class="text-sm text-[color:var(--text-muted)]">
            Maintained by the <strong>CloudySMS Core Team</strong>. <br>
            Empowering businesses with modern communication tools.
          </p>
          <p class="text-xs text-[color:var(--text-muted)] mt-2">
            Last Updated: August 4, 2026
          </p>
        </div>
        <div class="flex gap-6">
          <router-link to="/features" class="text-sm text-[color:var(--text-muted)] hover:text-[color:var(--text-main)]">Features</router-link>
          <router-link to="/pricing" class="text-sm text-[color:var(--text-muted)] hover:text-[color:var(--text-main)]">Pricing</router-link>
          <router-link to="/contact" class="text-sm text-[color:var(--text-muted)] hover:text-[color:var(--text-main)]">Contact Sales</router-link>
          <a href="https://github.com/cloudysms" target="_blank" rel="noopener noreferrer" class="text-sm text-[color:var(--text-muted)] hover:text-[color:var(--text-main)]">GitHub</a>
        </div>
      </div>
    </footer>

  </MarketingLayout>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap');

/* Base Styles & Variables */
/* Animations */
.fade-in-up {
  opacity: 0;
  transform: translateY(20px);
  animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.delay-1 { animation-delay: 0.1s; }
.delay-2 { animation-delay: 0.2s; }
.delay-3 { animation-delay: 0.3s; }
.delay-4 { animation-delay: 0.4s; }

@keyframes fadeInUp {
  to { opacity: 1; transform: translateY(0); }
}

/* Navbar */
.navbar {
  position: relative;
  z-index: 10;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.25rem 4rem;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--glass-border);
  background: rgba(255, 255, 255, 0.8);
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-weight: 700;
  font-size: 1.25rem;
  letter-spacing: -0.02em;
}

.logo-icon {
  width: 2.25rem;
  height: 2.25rem;
  border-radius: 0.6rem;
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-secondary));
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 15px rgba(16, 185, 129, 0.25);
}

.logo-icon.small {
  width: 1.75rem;
  height: 1.75rem;
  border-radius: 0.4rem;
}

.icon {
  width: 1.25rem;
  height: 1.25rem;
  color: var(--text-main);
}

.icon.small-icon {
  width: 1rem;
  height: 1rem;
}

.nav-links {
  display: flex;
  gap: 2.5rem;
}

.nav-links a {
  color: var(--text-muted);
  text-decoration: none;
  font-size: 0.95rem;
  font-weight: 500;
  transition: color 0.2s;
}

.nav-links a:hover {
  color: var(--text-main);
}

.nav-actions {
  display: flex;
  gap: 1rem;
}

/* Buttons */
.btn {
  padding: 0.5rem 1.25rem;
  border-radius: 0.5rem;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.btn-lg {
  padding: 0.875rem 2rem;
  font-size: 1rem;
  border-radius: 0.6rem;
}

.w-full {
  width: 100%;
}

.mt-auto {
  margin-top: auto;
}

.btn-primary {
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-secondary));
  color: white;
  box-shadow: 0 4px 15px rgba(16, 185, 129, 0.2);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(16, 185, 129, 0.35);
}

.pulse-hover:hover {
  animation: pulse-shadow 2s infinite;
}

@keyframes pulse-shadow {
  0% { box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.4); }
  70% { box-shadow: 0 0 0 10px rgba(16, 185, 129, 0); }
  100% { box-shadow: 0 0 0 0 rgba(16, 185, 129, 0); }
}

.btn-outline {
  background: transparent;
  color: var(--text-main);
  border: 1px solid var(--glass-border);
}

.btn-outline:hover {
  background: var(--bg-color-alt);
}

.btn-glass {
  background: white;
  color: var(--text-main);
  border: 1px solid var(--glass-border);
  box-shadow: 0 2px 5px rgba(0,0,0,0.05);
}

.btn-glass:hover {
  background: var(--bg-color-alt);
}

.arrow-icon {
  transition: transform 0.2s;
}

.btn-primary:hover .arrow-icon {
  transform: translateX(4px);
}

/* Hero Section */
.hero-section {
  position: relative;
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8rem 2rem 4rem;
  gap: 4rem;
  max-width: 1400px;
  margin: 0 auto;
}

.hero-section::before {
  content: '';
  position: absolute;
  top: -10%;
  left: 50%;
  transform: translateX(-50%);
  width: 100vw;
  height: 100%;
  background: radial-gradient(circle at top center, rgba(16, 185, 129, 0.1) 0%, rgba(14, 165, 233, 0.05) 30%, transparent 70%);
  z-index: -1;
  pointer-events: none;
}

.hero-content {
  flex: 1;
  text-align: left;
  max-width: 650px;
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.35rem 0.85rem;
  background: white;
  border: 1px solid var(--glass-border);
  border-radius: 2rem;
  color: var(--text-main);
  font-size: 0.85rem;
  font-weight: 500;
  margin-bottom: 2rem;
  box-shadow: 0 2px 5px rgba(0,0,0,0.05);
}

.badge-dot {
  width: 6px;
  height: 6px;
  background: #10b981;
  border-radius: 50%;
  box-shadow: 0 0 10px rgba(16, 185, 129, 0.4);
}

.hero-title {
  font-size: 5rem;
  line-height: 1.1;
  font-weight: 800;
  letter-spacing: -0.03em;
  margin-bottom: 1.5rem;
}

.text-gradient {
  background: linear-gradient(135deg, #0ea5e9, #10b981);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  color: transparent;
}

.hero-subtitle {
  font-size: 1.25rem;
  line-height: 1.6;
  color: var(--text-muted);
  margin-bottom: 2.5rem;
  max-width: 650px;
}

.hero-actions {
  display: flex;
  justify-content: flex-start;
  gap: 1rem;
}

/* Hero Visual (Glassmorphism Mockup) */
.hero-visual {
  flex: 1;
  width: 100%;
  max-width: 600px;
  perspective: 1200px;
}

.glass-panel {
  background: white;
  border: 1px solid var(--glass-border);
  border-radius: 1.25rem;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.mockup-panel {
  transform: rotateX(8deg) scale(0.95);
  transform-origin: bottom center;
  transition: transform 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  animation: float-mockup 6s infinite ease-in-out;
}

.mockup-panel:hover {
  transform: translateY(-10px) rotateX(0deg) scale(1);
}

@keyframes float-mockup {
  0%, 100% { transform: translateY(0) rotateX(8deg) scale(0.95); }
  50% { transform: translateY(-20px) rotateX(4deg) scale(0.97); }
}

.window-header {
  padding: 1rem;
  display: flex;
  gap: 0.5rem;
  border-bottom: 1px solid var(--glass-border);
  background: var(--bg-color);
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.dot.red { background: #ef4444; }
.dot.yellow { background: #eab308; }
.dot.green { background: #22c55e; }

.dashboard-mockup {
  display: flex;
  height: 400px;
  padding: 1rem;
  gap: 1rem;
}

.sidebar-mock {
  width: 60px;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.sidebar-item {
  height: 40px;
  border-radius: 0.5rem;
  background: var(--glass-border);
}

.sidebar-item.active {
  background: rgba(16, 185, 129, 0.2);
  border: 1px solid rgba(16, 185, 129, 0.3);
}

.main-mock {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.mock-header {
  height: 40px;
  border-radius: 0.5rem;
  background: var(--glass-border);
}

.mock-cards {
  display: flex;
  gap: 1rem;
}

.mock-card {
  flex: 1;
  height: 80px;
  border-radius: 0.5rem;
  background: var(--glass-border);
}

.card-animate-1 { animation: pulse-bg 3s infinite; }
.card-animate-2 { animation: pulse-bg 3s infinite 1s; }
.card-animate-3 { animation: pulse-bg 3s infinite 2s; }

@keyframes pulse-bg {
  0%, 100% { background: var(--glass-border); }
  50% { background: rgba(255, 255, 255, 0.12); }
}

.mock-chart {
  flex: 1;
  border-radius: 0.5rem;
  background: linear-gradient(to top, rgba(16, 185, 129, 0.05), transparent);
  border: 1px solid var(--glass-border);
  position: relative;
  overflow: hidden;
}

.mock-chart::after {
  content: '';
  position: absolute;
  bottom: 20%;
  left: 0;
  width: 100%;
  height: 2px;
  background: var(--accent-primary);
  box-shadow: 0 0 10px var(--accent-primary);
  transform: rotate(-5deg);
  animation: draw-line 3s infinite alternate ease-in-out;
}

@keyframes draw-line {
  0% { transform: rotate(-5deg) translateY(10px); opacity: 0.5;}
  100% { transform: rotate(-2deg) translateY(-10px); opacity: 1;}
}


/* Common Sections */
.section-container {
  padding: 6rem 4rem;
  max-width: 1400px;
  margin: 0 auto;
  position: relative;
  z-index: 10;
  width: 100%;
}

.alternate-bg {
  background-color: var(--bg-color-alt);
  max-width: none;
}

.alternate-bg > * {
  max-width: 1400px;
  margin: 0 auto;
}

.section-header {
  margin-bottom: 4rem;
}

.text-center {
  text-align: center;
}

.text-left {
  text-align: left;
}

.section-title {
  font-size: 2.75rem;
  font-weight: 700;
  margin-bottom: 1rem;
  letter-spacing: -0.02em;
}

.section-subtitle {
  font-size: 1.15rem;
  color: var(--text-muted);
  max-width: 600px;
  margin: 0 auto;
}

/* Features Section */
.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.feature-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 2rem;
  transition: all 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-5px);
  background: var(--bg-color-alt);
  border-color: #3f3f46;
}

.feature-icon-wrapper {
  width: 3rem;
  height: 3rem;
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

.feature-icon {
  width: 1.5rem;
  height: 1.5rem;
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

/* Solutions Section */
.solutions-content {
  display: flex;
  align-items: center;
  gap: 4rem;
}

.solutions-text {
  flex: 1;
}

.solutions-list {
  list-style: none;
  padding: 0;
  margin-top: 2rem;
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.solutions-list li {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  font-size: 1.05rem;
  color: var(--text-muted);
  line-height: 1.5;
}

.list-icon {
  color: var(--accent-primary);
  width: 1.5rem;
  height: 1.5rem;
  flex-shrink: 0;
  margin-top: 0.1rem;
}

.solutions-list strong {
  color: var(--text-main);
  font-weight: 600;
}

.solutions-visual {
  flex: 1;
}

.showcase-panel {
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.showcase-item {
  display: flex;
  gap: 1rem;
  align-items: center;
  padding: 1rem;
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 0.75rem;
}

.showcase-item.right {
  align-self: flex-end;
  background: rgba(16, 185, 129, 0.05);
  border-color: rgba(16, 185, 129, 0.2);
}

.showcase-icon {
  width: 2rem;
  height: 2rem;
  padding: 0.4rem;
  background: var(--glass-border);
  border-radius: 50%;
  color: var(--text-muted);
}

.user-avatar {
  width: 2rem;
  height: 2rem;
  background: var(--accent-primary);
  border-radius: 50%;
}

.showcase-lines {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.showcase-lines.right-align {
  align-items: flex-end;
}

.s-line {
  height: 6px;
  background: var(--glass-border);
  border-radius: 1rem;
}

.showcase-item.right .s-line {
  background: rgba(16, 185, 129, 0.3);
}

.w-full { width: 100%; }
.w-3\/4 { width: 75%; }
.w-1\/2 { width: 50%; }
.w-5\/6 { width: 83.33%; }


/* Pricing Section */
.pricing-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  align-items: stretch;
}

.pricing-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1.25rem;
  padding: 2.5rem 2rem;
  display: flex;
  flex-direction: column;
  position: relative;
  transition: transform 0.3s ease;
}

.pricing-card:hover {
  transform: translateY(-5px);
}

.pricing-card.popular {
  background: var(--bg-color-alt);
  border-color: var(--accent-primary);
  box-shadow: 0 10px 40px rgba(16, 185, 129, 0.1);
}

.popular-badge {
  position: absolute;
  top: -12px;
  left: 50%;
  transform: translateX(-50%);
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-secondary));
  color: white;
  padding: 0.25rem 1rem;
  border-radius: 1rem;
  font-size: 0.8rem;
  font-weight: 600;
  letter-spacing: 0.05em;
  text-transform: uppercase;
}

.plan-name {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1rem;
}

.plan-price {
  margin-bottom: 1rem;
  display: flex;
  align-items: flex-end;
}

.currency {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  color: var(--text-muted);
}

.amount {
  font-size: 3.5rem;
  font-weight: 800;
  line-height: 1;
  letter-spacing: -0.05em;
}

.period {
  font-size: 1rem;
  color: var(--text-muted);
  margin-bottom: 0.5rem;
  margin-left: 0.25rem;
}

.plan-desc {
  color: var(--text-muted);
  font-size: 0.95rem;
  margin-bottom: 2rem;
  line-height: 1.5;
}

.plan-features {
  list-style: none;
  padding: 0;
  margin: 0 0 2.5rem 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.plan-features li {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 0.95rem;
}

.check {
  color: var(--accent-primary);
  width: 1.25rem;
  height: 1.25rem;
  flex-shrink: 0;
}

/* Responsive */
@media (max-width: 1024px) {
  .hero-section, .solutions-content {
    flex-direction: column;
    text-align: center;
    padding: 4rem 2rem;
  }
  
  .hero-content, .solutions-text {
    margin: 0 auto;
  }
  
  .hero-actions {
    justify-content: center;
  }
  
  .nav-links {
    display: none;
  }
  
  .hero-visual, .solutions-visual {
    width: 100%;
    margin-top: 2rem;
  }
  
  .solutions-list li {
    justify-content: center;
    text-align: left;
  }
}

@media (max-width: 768px) {
  .hero-title {
    font-size: 3rem;
  }
  .section-container {
    padding: 4rem 2rem;
  }
  .footer-links {
    gap: 2rem;
  }
}
</style>
