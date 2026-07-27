<template>
  <div class="billing-view success-view">
    <div class="content-box fade-in-up">
      <div class="icon-wrapper green">
        <CheckCircle2 class="icon" />
      </div>
      <h1 class="title">Subscription Successful!</h1>
      <p class="subtitle">Thank you for subscribing to CloudySMS. Your account has been upgraded.</p>
      
      <div class="actions">
        <button @click="goToDashboard" class="btn btn-primary">
          Go to Dashboard
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { CheckCircle2 } from 'lucide-vue-next'

const router = useRouter()

const goToDashboard = () => {
  router.push('/dashboard')
}
</script>

<style scoped>
.billing-view {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: var(--surface-bg);
  padding: 2rem;
}

.content-box {
  background: white;
  padding: 3rem;
  border-radius: 16px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.05);
  text-align: center;
  max-width: 480px;
  width: 100%;
}

.icon-wrapper {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 1.5rem;
}

.icon-wrapper.green {
  background: rgba(34, 197, 94, 0.1);
  color: #22c55e;
}

.icon {
  width: 40px;
  height: 40px;
}

.title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 1rem;
}

.subtitle {
  font-size: 1.125rem;
  color: var(--text-secondary);
  margin-bottom: 2rem;
  line-height: 1.6;
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.875rem 2rem;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
}

.btn-primary {
  background: var(--primary-color);
  color: white;
}

.btn-primary:hover {
  background: var(--primary-dark);
  transform: translateY(-2px);
}
</style>
