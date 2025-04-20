<script setup lang="ts">
import { useAuthStore } from '~/stores/authStore';

const authStore = useAuthStore()
const router = useRouter()
const toast = useToast()


const handleLogout = async () => {
  await authStore.logout()
  toast.success({title: 'Logged out successfully'})
  router.push('/sign-in')
}
</script>

<template>
  <button
    class="text-white bg-red-500 hover:bg-red-600 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-4 py-2 text-center"
    @click="handleLogout"
  >
    Logout
  </button>
</template>