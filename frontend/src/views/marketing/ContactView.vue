<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'
import { Mail, MapPin, Phone } from 'lucide-vue-next'

useSeo({
  title: 'Contact Us - CloudySMS',
  description: 'Get in touch with the CloudySMS team for support, sales, or general inquiries.'
})
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Get in <span class="text-gradient">Touch</span></h1>
        <p class="page-subtitle fade-in-up delay-1">Have a question or want to learn more about our enterprise plans? We'd love to hear from you.</p>
      </div>

      <div class="contact-content fade-in-up delay-2">
        <div class="contact-info">
          <div class="info-card">
            <div class="icon-wrapper">
              <Mail class="icon" />
            </div>
            <h3>Email Us</h3>
            <p>Our friendly team is here to help.</p>
            <a href="mailto:hello@cloudysms.com">hello@cloudysms.com</a>
          </div>
          <div class="info-card">
            <div class="icon-wrapper">
              <MapPin class="icon" />
            </div>
            <h3>Visit Us</h3>
            <p>Come say hello at our office HQ.</p>
            <span>100 Innovation Drive, San Francisco, CA 94103</span>
          </div>
          <div class="info-card">
            <div class="icon-wrapper">
              <Phone class="icon" />
            </div>
            <h3>Call Us</h3>
            <p>Mon-Fri from 8am to 5pm.</p>
            <a href="tel:+15550000000">+1 (555) 000-0000</a>
          </div>
        </div>

        <div class="contact-form-wrapper">
          <form class="contact-form" @submit.prevent>
            <h2>Send us a message</h2>
            <div class="form-group">
              <label for="name">Name</label>
              <input type="text" id="name" placeholder="Jane Doe" required />
            </div>
            <div class="form-group">
              <label for="email">Email</label>
              <input type="email" id="email" placeholder="jane@example.com" required />
            </div>
            <div class="form-group">
              <label for="subject">Subject</label>
              <select id="subject" required>
                <option value="">Select a topic...</option>
                <option value="sales">Sales & Pricing</option>
                <option value="support">Technical Support</option>
                <option value="partnership">Partnerships</option>
                <option value="other">Other</option>
              </select>
            </div>
            <div class="form-group">
              <label for="message">Message</label>
              <textarea id="message" rows="5" placeholder="How can we help?" required></textarea>
            </div>
            <button type="submit" class="btn btn-primary w-full">Send Message</button>
          </form>
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

.contact-content {
  display: flex;
  gap: 4rem;
  align-items: flex-start;
}

.contact-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.info-card {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 2rem;
}

.icon-wrapper {
  width: 3rem;
  height: 3rem;
  border-radius: 0.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(16, 185, 129, 0.1);
  color: #10b981;
  margin-bottom: 1rem;
}

.info-card h3 {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.info-card p {
  color: var(--text-muted);
  font-size: 0.95rem;
  margin-bottom: 1rem;
}

.info-card a, .info-card span {
  color: #10b981;
  font-weight: 500;
  text-decoration: none;
}

.info-card a:hover {
  text-decoration: underline;
}

.contact-form-wrapper {
  flex: 1.2;
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1.5rem;
  padding: 3rem;
  backdrop-filter: blur(10px);
}

.contact-form h2 {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  font-size: 0.9rem;
  font-weight: 500;
  margin-bottom: 0.5rem;
  color: var(--text-main);
}

input, select, textarea {
  width: 100%;
  padding: 0.875rem 1rem;
  border-radius: 0.5rem;
  background: var(--bg-color);
  border: 1px solid var(--glass-border);
  color: var(--text-main);
  font-family: inherit;
  font-size: 1rem;
  transition: border-color 0.2s, box-shadow 0.2s;
  outline: none;
}

input:focus, select:focus, textarea:focus {
  border-color: #10b981;
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
}

select {
  appearance: none;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='white' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3e%3cpolyline points='6 9 12 15 18 9'%3e%3c/polyline%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: right 1rem center;
  background-size: 1em;
}

select option {
  background-color: #111113;
  color: var(--text-main);
}

.btn {
  padding: 1rem 2rem;
  border-radius: 0.5rem;
  font-size: 1rem;
  font-weight: 600;
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

.w-full {
  width: 100%;
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
  .contact-content { flex-direction: column; }
  .contact-form-wrapper { width: 100%; }
}

@media (max-width: 768px) {
  .page-container { padding: 4rem 2rem; }
  .page-title { font-size: 2.5rem; }
  .contact-form-wrapper { padding: 2rem; }
}
</style>
