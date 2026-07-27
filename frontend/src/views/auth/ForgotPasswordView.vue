<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { toast } from 'vue-sonner'
import { MessageSquare, Loader2, Mail, ArrowLeft } from 'lucide-vue-next'
import { useSeo } from '@/composables/useSeo'
import api from '@/services/api'

const { t } = useI18n()

useSeo({
  title: 'Forgot Password',
  description: 'Reset your CloudySMS password.'
})

const router = useRouter()
const email = ref('')
const isLoading = ref(false)
const isSent = ref(false)

const handleForgotPassword = async () => {
  if (!email.value) {
    toast.error('Please enter your email')
    return
  }

  isLoading.value = true

  try {
    await api.post('/auth/forgot-password', { email: email.value })
    isSent.value = true
    toast.success('Reset link sent!')
  } catch (error: any) {
    const message = error.response?.data?.message || 'Failed to request reset'
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
        <h2 class="auth-title">Forgot Password</h2>
        <p class="auth-subtitle">
          {{ isSent ? 'Check your email' : 'Enter your email to receive a reset link' }}
        </p>
      </div>

      <div v-if="isSent" class="sent-message text-center">
        <p class="text-sm text-gray-400 mb-6">
          If an account exists for <strong>{{ email }}</strong>, you will receive a password reset link shortly. Please check your spam folder as well.
        </p>
        <RouterLink to="/login" class="btn btn-primary btn-block">
          Return to login
        </RouterLink>
      </div>

      <form v-else @submit.prevent="handleForgotPassword" class="auth-form">
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

        <button type="submit" class="btn btn-primary btn-block pulse-hover mt-2" :disabled="isLoading">
          <Loader2 v-if="isLoading" class="spinner" />
          <span v-else>Send Reset Link</span>
        </button>
      </form>

      <div class="auth-footer" v-if="!isSent">
        <RouterLink to="/login" class="back-link">
          <ArrowLeft class="w-4 h-4 mr-1" /> Back to login
        </RouterLink>
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
  background-image: radial-gradient(circle at top left, rgba(16, 185, 129, 0.1), transparent 40%),
                    radial-gradient(circle at bottom right, rgba(16, 185, 129, 0.05), transparent 40%);
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

.auth-footer {
  margin-top: 2rem;
  text-align: center;
}

.back-link {
  display: inline-flex;
  align-items: center;
  font-size: 0.875rem;
  color: var(--text-muted);
  text-decoration: none;
  transition: color 0.2s ease;
}

.back-link:hover {
  color: var(--primary-light);
}
</style>
