<script setup lang="ts">
import { useRouter } from 'vue-router'
import { MessageSquare } from 'lucide-vue-next'

const router = useRouter()

const getStarted = () => {
  router.push('/login')
}
</script>

<template>
  <div class="marketing-layout">

    <header class="navbar">
      <div class="logo cursor-pointer" @click="router.push('/')">
        <div class="logo-icon">
          <MessageSquare class="icon" />
        </div>
        <span class="logo-text">CloudySMS</span>
      </div>
      <nav class="nav-links">
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/features">Features</RouterLink>
        <RouterLink to="/pricing">Pricing</RouterLink>
        <RouterLink to="/integrations">Integrations</RouterLink>
      </nav>
      <div class="nav-actions">
        <button class="btn btn-outline" @click="router.push('/login')">Log In</button>
        <button class="btn btn-primary" @click="getStarted">Get Started</button>
      </div>
    </header>

    <div class="content-wrapper">
      <slot />
    </div>

    <!-- Footer -->
    <footer class="footer">
      <div class="footer-content">
        <div class="footer-brand">
          <div class="logo">
            <div class="logo-icon small">
              <MessageSquare class="icon small-icon" />
            </div>
            <span class="logo-text">CloudySMS</span>
          </div>
          <p class="footer-desc">The all-in-one WhatsApp Business Platform for modern teams.</p>
        </div>
        <div class="footer-links">
          <div class="link-group">
            <h4>Product</h4>
            <RouterLink to="/features">Features</RouterLink>
            <RouterLink to="/pricing">Pricing</RouterLink>
            <RouterLink to="/integrations">Integrations</RouterLink>
            <RouterLink to="/changelog">Changelog</RouterLink>
          </div>
          <div class="link-group">
            <h4>Resources</h4>
            <RouterLink to="/docs">Documentation</RouterLink>
            <RouterLink to="/api-reference">API Reference</RouterLink>
            <RouterLink to="/blog">Blog</RouterLink>
            <RouterLink to="/community">Community</RouterLink>
          </div>
          <div class="link-group">
            <h4>Company</h4>
            <RouterLink to="/about">About</RouterLink>
            <RouterLink to="/contact">Contact</RouterLink>
            <RouterLink to="/privacy">Privacy Policy</RouterLink>
            <RouterLink to="/terms">Terms of Service</RouterLink>
          </div>
        </div>
      </div>
      <div class="footer-bottom">
        <p>&copy; 2026 CloudySMS. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap');

/* Base Styles & Variables */
.marketing-layout {
  --bg-color: #ffffff;
  --bg-color-alt: #f8fafc;
  --text-main: #0f172a;
  --text-muted: #64748b;
  --accent-primary: #0f172a;
  --accent-secondary: #334155;
  --glass-bg: #ffffff;
  --glass-border: #e2e8f0;
  
  font-family: 'Inter', sans-serif;
  background-color: var(--bg-color);
  color: var(--text-main);
  min-height: 100vh;
  position: relative;
  overflow-x: hidden;
  display: flex;
  flex-direction: column;
}

.cursor-pointer {
  cursor: pointer;
}

.content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
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
  border-radius: 6px;
  background: var(--text-main);
  color: var(--bg-color);
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo-icon.small {
  width: 1.75rem;
  height: 1.75rem;
  border-radius: 4px;
}

.icon {
  width: 1.25rem;
  height: 1.25rem;
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

.nav-links a:hover,
.nav-links a.router-link-exact-active {
  color: var(--text-main);
}

.nav-actions {
  display: flex;
  gap: 1rem;
}

/* Buttons */
.btn {
  padding: 0.5rem 1.25rem;
  border-radius: 6px;
  font-size: 0.85rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  border: 1px solid transparent;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  letter-spacing: 0.01em;
}

.btn-primary {
  background: var(--text-main);
  color: var(--bg-color);
}

.btn-primary:hover {
  background: #334155;
}

.btn-outline {
  background: transparent;
  color: var(--text-main);
  border: 1px solid var(--glass-border);
}

.btn-outline:hover {
  background: var(--glass-border);
}

/* Footer */
.footer {
  border-top: 1px solid var(--glass-border);
  background: var(--bg-color-alt);
  padding: 4rem 4rem 2rem;
  margin-top: auto;
}

.footer-content {
  max-width: 1400px;
  margin: 0 auto;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  gap: 3rem;
  margin-bottom: 3rem;
}

.footer-brand {
  flex: 1;
  min-width: 250px;
  max-width: 350px;
}

.footer-brand .logo {
  margin-bottom: 1rem;
}

.footer-desc {
  color: var(--text-muted);
  font-size: 0.95rem;
  line-height: 1.6;
}

.footer-links {
  display: flex;
  gap: 4rem;
  flex-wrap: wrap;
}

.link-group {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.link-group h4 {
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--text-main);
  margin-bottom: 0.5rem;
}

.link-group a {
  color: var(--text-muted);
  text-decoration: none;
  font-size: 0.9rem;
  transition: color 0.2s;
}

.link-group a:hover {
  color: var(--text-main);
}

.footer-bottom {
  max-width: 1400px;
  margin: 0 auto;
  padding-top: 2rem;
  border-top: 1px solid var(--glass-border);
  text-align: center;
  color: var(--text-muted);
  font-size: 0.85rem;
}

/* Responsive */
@media (max-width: 1024px) {
  .nav-links {
    display: none;
  }
}
@media (max-width: 768px) {
  .footer-links {
    gap: 2rem;
  }
}
</style>
