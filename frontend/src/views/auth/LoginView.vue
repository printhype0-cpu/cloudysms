<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAuthStore } from '@/stores/auth'
import { api } from '@/services/api'
import { toast } from 'vue-sonner'
import { MessageSquare, Loader2, Mail, Lock } from 'lucide-vue-next'
import { useSeo } from '@/composables/useSeo'

const { t } = useI18n()

useSeo({
  title: 'Log In',
  description: 'Access your CloudySMS account to manage conversations and automated workflows.'
})

interface SSOProvider {
  provider: string
  name: string
}

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const isLoading = ref(false)
const ssoProviders = ref<SSOProvider[]>([])

// SSO provider icons (using simple SVG paths)
const providerIcons: Record<string, string> = {
  google: 'M12.545,10.239v3.821h5.445c-0.712,2.315-2.647,3.972-5.445,3.972c-3.332,0-6.033-2.701-6.033-6.032s2.701-6.032,6.033-6.032c1.498,0,2.866,0.549,3.921,1.453l2.814-2.814C17.503,2.988,15.139,2,12.545,2C7.021,2,2.543,6.477,2.543,12s4.478,10,10.002,10c8.396,0,10.249-7.85,9.426-11.748L12.545,10.239z',
  microsoft: 'M11 11H3V3h8v8zm10 0h-8V3h8v8zM11 21H3v-8h8v8zm10 0h-8v-8h8v8z',
  github: 'M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z',
  facebook: 'M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z',
  custom: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm-1-13h2v6h-2zm0 8h2v2h-2z'
}

onMounted(async () => {
  // Check for SSO error in query params
  const ssoError = route.query.sso_error as string
  if (ssoError) {
    toast.error(decodeURIComponent(ssoError))
    router.replace({ query: { ...route.query, sso_error: undefined } })
  }

  // Fetch enabled SSO providers
  try {
    const response = await api.get('/auth/sso/providers')
    ssoProviders.value = response.data.data || []
  } catch {
    ssoProviders.value = []
  }
})

const handleLogin = async () => {
  if (!email.value || !password.value) {
    toast.error(t('auth.enterEmailPassword'))
    return
  }

  isLoading.value = true

  try {
    await authStore.login(email.value, password.value)
    toast.success(t('auth.loginSuccess'))

    const redirect = route.query.redirect as string
    router.push(redirect || '/app/dashboard')
  } catch (error: any) {
    const message = error.response?.data?.message || t('auth.invalidCredentials')
    toast.error(message)
  } finally {
    isLoading.value = false
  }
}

const initiateSSO = (provider: string) => {
  const basePath = ((window as any).__BASE_PATH__ ?? '').replace(/\/$/, '')
  window.location.href = `${basePath}/api/auth/sso/${provider}/init`
}
</script>

<template>
  <div class="auth-layout">
    <div class="auth-card fade-in-up">
      <div class="auth-header">
        <div class="logo-wrapper">
          <MessageSquare class="logo-icon" />
        </div>
        <h2 class="auth-title">{{ $t('auth.welcomeTitle') }}</h2>
        <p class="auth-subtitle">{{ $t('auth.welcomeSubtitle') }}</p>
      </div>

      <form @submit.prevent="handleLogin" class="auth-form">
        <div class="form-group">
          <label for="email" class="form-label">{{ $t('common.email') }}</label>
          <div class="input-wrapper">
            <Mail class="input-icon" />
            <input
              id="email"
              v-model="email"
              type="email"
              class="form-input"
              :placeholder="$t('auth.emailPlaceholder')"
              :disabled="isLoading"
              autocomplete="email"
            />
          </div>
        </div>
        <div class="form-group">
          <label for="password" class="form-label">{{ $t('auth.password') }}</label>
          <div class="input-wrapper">
            <Lock class="input-icon" />
            <input
              id="password"
              v-model="password"
              type="password"
              class="form-input"
              :placeholder="$t('auth.passwordPlaceholder')"
              :disabled="isLoading"
              autocomplete="current-password"
            />
          </div>
        </div>
        
        <button type="submit" class="btn btn-primary btn-block pulse-hover" :disabled="isLoading">
          <Loader2 v-if="isLoading" class="spinner" />
          <span v-else>{{ $t('auth.signIn') }}</span>
        </button>
      </form>

      <!-- SSO Section -->
      <div v-if="ssoProviders.length > 0" class="sso-section">
        <div class="divider">
          <span>{{ $t('auth.orContinueWith') }}</span>
        </div>

        <button
          v-for="provider in ssoProviders"
          :key="provider.provider"
          class="btn btn-glass btn-block sso-btn"
          @click="initiateSSO(provider.provider)"
        >
          <svg class="sso-icon" viewBox="0 0 24 24" fill="currentColor">
            <path :d="providerIcons[provider.provider] || providerIcons.custom" />
          </svg>
          {{ provider.name }}
        </button>
      </div>

      <div class="auth-footer">
        <p class="footer-text">
          {{ $t('auth.noAccount') }}
          <RouterLink to="/register" class="footer-link">{{ $t('auth.signUp') }}</RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.auth-layout {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--bg-color);
  background-image: radial-gradient(circle at top right, rgba(16, 185, 129, 0.1), transparent 40%),
                    radial-gradient(circle at bottom left, rgba(16, 185, 129, 0.05), transparent 40%);
  padding: 1rem;
}

.auth-card {
  width: 100%;
  max-width: 440px;
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1.5rem;
  padding: 2.5rem 2rem;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  backdrop-filter: blur(12px);
}

.auth-header {
  text-align: center;
  margin-bottom: 2rem;
}

.logo-wrapper {
  width: 3.5rem;
  height: 3.5rem;
  border-radius: 1rem;
  background: linear-gradient(135deg, var(--primary-color), var(--primary-dark));
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 1.25rem;
  box-shadow: 0 8px 24px rgba(16, 185, 129, 0.3);
}

.logo-icon {
  width: 1.75rem;
  height: 1.75rem;
  color: white;
}

.auth-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--text-main);
  margin-bottom: 0.5rem;
}

.auth-subtitle {
  font-size: 1rem;
  color: var(--text-muted);
}

.auth-form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-label {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-muted);
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.input-icon {
  position: absolute;
  left: 1rem;
  width: 1.25rem;
  height: 1.25rem;
  color: var(--text-muted);
  pointer-events: none;
}

.form-input {
  width: 100%;
  padding: 0.875rem 1rem 0.875rem 3rem;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid var(--glass-border);
  border-radius: 0.75rem;
  color: var(--text-main);
  font-size: 1rem;
  transition: all 0.2s ease;
}

.form-input:focus {
  outline: none;
  border-color: var(--primary-color);
  background: rgba(255, 255, 255, 0.05);
  box-shadow: 0 0 0 4px rgba(16, 185, 129, 0.1);
}

.form-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-block {
  width: 100%;
  margin-top: 0.5rem;
}

.spinner {
  animation: spin 1s linear infinite;
  width: 1.25rem;
  height: 1.25rem;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.sso-section {
  margin-top: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.divider {
  display: flex;
  align-items: center;
  text-align: center;
  margin: 0.5rem 0;
}

.divider::before,
.divider::after {
  content: '';
  flex: 1;
  border-bottom: 1px solid var(--glass-border);
}

.divider span {
  padding: 0 1rem;
  color: var(--text-muted);
  font-size: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  font-weight: 600;
}

.sso-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 0.875rem;
  border-radius: 0.75rem;
}

.sso-icon {
  width: 1.25rem;
  height: 1.25rem;
}

.auth-footer {
  margin-top: 2rem;
  text-align: center;
}

.footer-text {
  font-size: 0.875rem;
  color: var(--text-muted);
}

.footer-link {
  color: var(--primary-color);
  font-weight: 600;
  text-decoration: none;
  margin-left: 0.25rem;
  transition: color 0.2s ease;
}

.footer-link:hover {
  color: var(--primary-light);
  text-decoration: underline;
}
</style>
