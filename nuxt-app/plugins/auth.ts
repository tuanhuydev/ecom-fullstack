import { useAuthStore } from '~/stores/authStore'

export default defineNuxtPlugin(async () => {
  // Get the auth store
  const authStore = useAuthStore()
  
  // Initialize auth state - this will check for tokens and fetch user data if needed
  await authStore.initAuth()
})