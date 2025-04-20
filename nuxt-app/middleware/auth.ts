import { useAuthStore } from '~/stores/authStore'

export default defineNuxtRouteMiddleware(async (to) => {
  const authStore = useAuthStore()
  
  // Check if we have a token stored
  if (!authStore.token && typeof localStorage !== 'undefined') {
    const token = localStorage.getItem('auth_token')
    if (token) {
      authStore.token = token
      // Try to fetch user info with the token
      await authStore.fetchUserWithToken()
    }
  }
  
  // If user is still not authenticated after checking token
  if (!authStore.isAuthenticated) {
    return navigateTo('/sign-in')
  }
})