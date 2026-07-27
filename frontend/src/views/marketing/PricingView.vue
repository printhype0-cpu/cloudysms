<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'
import { CheckCircle2, X } from 'lucide-vue-next'

useSeo({
  title: 'Pricing - CloudySMS',
  description: 'Simple, transparent pricing. Start for free and upgrade when you need more power.'
})
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Simple, transparent <span class="text-gradient">pricing</span></h1>
        <p class="page-subtitle fade-in-up delay-1">Start for free, upgrade when you need more power.</p>
      </div>

      <div class="pricing-grid fade-in-up delay-2">
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
            <li class="disabled"><X class="x-icon" /> Custom API Access</li>
            <li class="disabled"><X class="x-icon" /> Role-based Access</li>
          </ul>
          <button class="btn btn-outline w-full mt-auto">Get Started</button>
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
            <li><CheckCircle2 class="check" /> Custom API Access</li>
            <li><CheckCircle2 class="check" /> Role-based Access</li>
            <li><CheckCircle2 class="check" /> Priority Support</li>
          </ul>
          <button class="btn btn-primary w-full mt-auto">Start Free Trial</button>
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
            <li><CheckCircle2 class="check" /> Custom Billing</li>
          </ul>
          <button class="btn btn-outline w-full mt-auto">Contact Sales</button>
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
  border-color: #10b981;
  box-shadow: 0 10px 40px rgba(16, 185, 129, 0.1);
}

.popular-badge {
  position: absolute;
  top: -12px;
  left: 50%;
  transform: translateX(-50%);
  background: linear-gradient(135deg, #10b981, #059669);
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

.plan-features li.disabled {
  color: #52525b;
}

.check {
  color: #10b981;
  width: 1.25rem;
  height: 1.25rem;
  flex-shrink: 0;
}

.x-icon {
  color: #52525b;
  width: 1.25rem;
  height: 1.25rem;
  flex-shrink: 0;
}

.btn {
  padding: 0.875rem 2rem;
  border-radius: 0.6rem;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.btn-primary {
  background: linear-gradient(135deg, #10b981, #059669);
  color: white;
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

.w-full { width: 100%; }
.mt-auto { margin-top: auto; }

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
