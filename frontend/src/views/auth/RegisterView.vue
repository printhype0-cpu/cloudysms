<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAuthStore } from '@/stores/auth'
import { toast } from 'vue-sonner'
import { MessageSquare, Loader2, Mail, Lock, User } from 'lucide-vue-next'
import { useSeo } from '@/composables/useSeo'

const { t } = useI18n()

useSeo({
  title: 'Sign Up',
  description: 'Create a new CloudySMS account and start automating your WhatsApp communication.'
})
const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const fullName = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const isLoading = ref(false)

const organizationId = computed(() => (route.query.org as string) || '')

const handleRegister = async () => {
  if (!organizationId.value) {
    toast.error(t('auth.invitationRequired'))
    return
  }

  if (!fullName.value || !email.value || !password.value) {
    toast.error(t('auth.fillAllFields'))
    return
  }

  if (password.value !== confirmPassword.value) {
    toast.error(t('auth.passwordsMismatch'))
    return
  }

  if (password.value.length < 8) {
    toast.error(t('auth.passwordTooShort'))
    return
  }

  isLoading.value = true

  try {
    await authStore.register({
      full_name: fullName.value,
      email: email.value,
      password: password.value,
      organization_id: organizationId.value
    })
    toast.success(t('auth.registrationSuccess'))
    router.push('/app/dashboard')
  } catch (error: any) {
    const message = error.response?.data?.message || t('auth.registrationFailed')
    toast.error(message)
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="auth-layout">
    <div class="auth-card fade-in-up">
      <div class="auth-header">
        <div class="logo-wrapper">
          <MessageSquare class="logo-icon" />
        </div>
        <h2 class="auth-title">{{ $t('auth.createAccount') }}</h2>
        <p class="auth-subtitle">{{ $t('auth.createAccountDesc') }}</p>
      </div>

      <!-- No org ID in URL — show invitation required message -->
      <template v-if="!organizationId">
        <div class="invitation-required">
          <p class="required-text">{{ $t('auth.invitationRequired') }}</p>
          <RouterLink to="/login" class="btn btn-outline btn-block mt-4">
            {{ $t('auth.signIn') }}
          </RouterLink>
        </div>
      </template>

      <!-- Has org ID — show registration form -->
      <form v-else @submit.prevent="handleRegister" class="auth-form">
        <div class="form-group">
          <label for="fullName" class="form-label">{{ $t('auth.fullName') }}</label>
          <div class="input-wrapper">
            <User class="input-icon" />
            <input
              id="fullName"
              v-model="fullName"
              type="text"
              class="form-input"
              :placeholder="$t('auth.fullNamePlaceholder')"
              :disabled="isLoading"
              autocomplete="name"
            />
          </div>
        </div>

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
              :placeholder="$t('auth.passwordMinLength')"
              :disabled="isLoading"
              autocomplete="new-password"
            />
          </div>
        </div>

        <div class="form-group">
          <label for="confirmPassword" class="form-label">{{ $t('auth.confirmPassword') }}</label>
          <div class="input-wrapper">
            <Lock class="input-icon" />
            <input
              id="confirmPassword"
              v-model="confirmPassword"
              type="password"
              class="form-input"
              :placeholder="$t('auth.confirmPasswordPlaceholder')"
              :disabled="isLoading"
              autocomplete="new-password"
            />
          </div>
        </div>
        
        <button type="submit" class="btn btn-primary btn-block pulse-hover" :disabled="isLoading">
          <Loader2 v-if="isLoading" class="spinner" />
          <span v-else>{{ $t('auth.createAccountBtn') }}</span>
        </button>
      </form>

      <div class="auth-footer" v-if="organizationId">
        <p class="footer-text">
          {{ $t('auth.alreadyHaveAccount') }}
          <RouterLink to="/login" class="footer-link">{{ $t('auth.signIn') }}</RouterLink>
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

.invitation-required {
  text-align: center;
  padding: 1.5rem 0;
}

.required-text {
  color: var(--text-muted);
  font-size: 0.95rem;
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

.mt-4 {
  margin-top: 1rem;
}
</style>
