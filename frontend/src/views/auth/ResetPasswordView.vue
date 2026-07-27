<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { toast } from 'vue-sonner'
import { MessageSquare, Loader2, Lock } from 'lucide-vue-next'
import { useSeo } from '@/composables/useSeo'
import api from '@/services/api'

const { t } = useI18n()

useSeo({
  title: 'Reset Password',
  description: 'Enter your new password for CloudySMS.'
})

const router = useRouter()
const route = useRoute()
const token = ref('')
const password = ref('')
const confirmPassword = ref('')
const isLoading = ref(false)

onMounted(() => {
  if (route.query.token) {
    token.value = route.query.token as string
  } else {
    toast.error('Invalid or missing password reset token.')
    router.push('/login')
  }
})

const handleResetPassword = async () => {
  if (!password.value || !confirmPassword.value) {
    toast.error('Please fill all fields')
    return
  }
  
  if (password.value !== confirmPassword.value) {
    toast.error('Passwords do not match')
    return
  }
  
  if (password.value.length < 8) {
    toast.error('Password must be at least 8 characters')
    return
  }

  isLoading.value = true

  try {
    await api.post('/auth/reset-password', { 
      token: token.value,
      password: password.value 
    })
    toast.success('Password successfully reset! Please log in.')
    router.push('/login')
  } catch (error: any) {
    const message = error.response?.data?.message || 'Failed to reset password'
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
        <h2 class="auth-title">Set New Password</h2>
        <p class="auth-subtitle">Please enter your new password</p>
      </div>

      <form @submit.prevent="handleResetPassword" class="auth-form">
        <div class="form-group">
          <label for="password" class="form-label">New Password</label>
          <div class="input-wrapper">
            <Lock class="input-icon" />
            <input
              id="password"
              v-model="password"
              type="password"
              class="form-input"
              placeholder="At least 8 characters"
              :disabled="isLoading"
              autocomplete="new-password"
            />
          </div>
        </div>

        <div class="form-group">
          <label for="confirmPassword" class="form-label">Confirm New Password</label>
          <div class="input-wrapper">
            <Lock class="input-icon" />
            <input
              id="confirmPassword"
              v-model="confirmPassword"
              type="password"
              class="form-input"
              placeholder="Repeat your new password"
              :disabled="isLoading"
              autocomplete="new-password"
            />
          </div>
        </div>

        <button type="submit" class="btn btn-primary btn-block pulse-hover mt-2" :disabled="isLoading">
          <Loader2 v-if="isLoading" class="spinner" />
          <span v-else>Update Password</span>
        </button>
      </form>
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
</style>
