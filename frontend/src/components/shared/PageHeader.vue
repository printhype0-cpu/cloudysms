<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb'
import { ArrowLeft } from 'lucide-vue-next'
import type { Component } from 'vue'

defineProps<{
  title: string
  description?: string
  icon?: Component
  iconGradient?: string
  backLink?: string
  breadcrumbs?: Array<{ label: string; href?: string }>
}>()
</script>

<template>
  <header class="border-b border-[var(--glass-border)] bg-[var(--bg-color)]/95 backdrop-blur">
    <div class="flex h-16 items-center px-6">
      <RouterLink v-if="backLink" :to="backLink">
        <Button variant="ghost" size="icon" class="mr-3 text-[var(--text-muted)] hover:text-[var(--text-main)] hover:bg-[var(--glass-bg)]">
          <ArrowLeft class="h-5 w-5" />
        </Button>
      </RouterLink>
      <div
        v-if="icon"
        class="h-8 w-8 rounded-lg flex items-center justify-center mr-3 shadow-lg"
        :class="iconGradient || 'bg-gradient-to-br from-[var(--primary-color)] to-[var(--primary-dark)] shadow-[var(--primary-color)]/20'"
      >
        <component :is="icon" class="h-4 w-4 text-white" />
      </div>
      <div class="flex-1">
        <h1 class="text-xl font-semibold text-[var(--text-main)]">{{ title }}</h1>
        <template v-if="breadcrumbs?.length">
          <Breadcrumb>
            <BreadcrumbList>
              <template v-for="(crumb, index) in breadcrumbs" :key="index">
                <BreadcrumbItem>
                  <BreadcrumbLink v-if="crumb.href" :href="crumb.href" class="text-[var(--text-muted)] hover:text-[var(--text-main)]">
                    {{ crumb.label }}
                  </BreadcrumbLink>
                  <BreadcrumbPage v-else class="text-[var(--text-main)]">{{ crumb.label }}</BreadcrumbPage>
                </BreadcrumbItem>
                <BreadcrumbSeparator v-if="index < breadcrumbs.length - 1" class="text-[var(--text-muted)]" />
              </template>
            </BreadcrumbList>
          </Breadcrumb>
        </template>
        <p v-else-if="description" class="text-sm text-[var(--text-muted)]">
          {{ description }}
        </p>
      </div>
      <slot name="actions" />
    </div>
  </header>
</template>
