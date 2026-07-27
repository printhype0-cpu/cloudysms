<script setup lang="ts">
import { useSeo } from '@/composables/useSeo'
import MarketingLayout from '@/components/layout/MarketingLayout.vue'

useSeo({
  title: 'Terms of Service - CloudySMS',
  description: 'Read the terms and conditions for using the CloudySMS platform.'
})
</script>

<template>
  <MarketingLayout>
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title fade-in-up">Terms of <span class="text-gradient">Service</span></h1>
        <p class="page-subtitle fade-in-up delay-1">Last updated: August 1, 2026</p>
      </div>

      <div class="legal-content fade-in-up delay-2">
        <section>
          <h2>1. Agreement to Terms</h2>
          <p>By accessing or using our services, you agree to be bound by these Terms of Service. If you disagree with any part of the terms, then you do not have permission to access the Service.</p>
        </section>

        <section>
          <h2>2. Use License</h2>
          <p>Permission is granted to temporarily download one copy of the materials (information or software) on CloudySMS's website for personal, non-commercial transitory viewing only. This is the grant of a license, not a transfer of title, and under this license you may not:</p>
          <ul>
            <li>modify or copy the materials;</li>
            <li>use the materials for any commercial purpose, or for any public display (commercial or non-commercial);</li>
            <li>attempt to decompile or reverse engineer any software contained on CloudySMS's website;</li>
            <li>remove any copyright or other proprietary notations from the materials; or</li>
            <li>transfer the materials to another person or "mirror" the materials on any other server.</li>
          </ul>
        </section>

        <section>
          <h2>3. Disclaimer</h2>
          <p>The materials on CloudySMS's website are provided on an 'as is' basis. CloudySMS makes no warranties, expressed or implied, and hereby disclaims and negates all other warranties including, without limitation, implied warranties or conditions of merchantability, fitness for a particular purpose, or non-infringement of intellectual property or other violation of rights.</p>
        </section>

        <section>
          <h2>4. Limitations</h2>
          <p>In no event shall CloudySMS or its suppliers be liable for any damages (including, without limitation, damages for loss of data or profit, or due to business interruption) arising out of the use or inability to use the materials on CloudySMS's website, even if CloudySMS or a CloudySMS authorized representative has been notified orally or in writing of the possibility of such damage.</p>
        </section>

        <section>
          <h2>5. Governing Law</h2>
          <p>These terms and conditions are governed by and construed in accordance with the laws and you irrevocably submit to the exclusive jurisdiction of the courts in that State or location.</p>
        </section>
      </div>
    </div>
  </MarketingLayout>
</template>

<style scoped>
.page-container {
  padding: 6rem 4rem;
  max-width: 900px;
  margin: 0 auto;
  width: 100%;
}

.page-header {
  text-align: left;
  margin-bottom: 4rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 2rem;
}

.page-title {
  font-size: 3.5rem;
  font-weight: 800;
  margin-bottom: 1rem;
  letter-spacing: -0.03em;
}

.text-gradient {
  color: var(--text-main);
}

.page-subtitle {
  font-size: 1.1rem;
  color: var(--text-muted);
}

.legal-content {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  border-radius: 1rem;
  padding: 3rem;
}

section {
  margin-bottom: 3rem;
}

section:last-child {
  margin-bottom: 0;
}

h2 {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1.25rem;
  color: var(--text-main);
}

p {
  color: var(--text-muted);
  line-height: 1.7;
  font-size: 1.05rem;
  margin-bottom: 1.25rem;
}

ul {
  color: var(--text-muted);
  line-height: 1.7;
  font-size: 1.05rem;
  margin-bottom: 1.25rem;
  padding-left: 1.5rem;
}

li {
  margin-bottom: 0.5rem;
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
  .legal-content { padding: 2rem; }
}
</style>
