<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod';
import { useForm } from 'vee-validate';
import { z } from 'zod';
import InputField from '~/components/InputField.vue';
import { useAuthStore } from '~/stores/authStore';

// Using Nuxt's built-in composable for toast
const router = useRouter();
const authStore = useAuthStore();
const toast = useToast()

const signInSchema = z.object({
  email: z.string().email('Invalid email address'),
  password: z.string().min(6, 'Password must be at least 6 characters long'),
});

const { handleSubmit, resetForm } = useForm({
  validationSchema: toTypedSchema(signInSchema),
  initialValues: {
    email: '',
    password: '',
  },
});

const onSubmit = handleSubmit(async (values) => {
  try {
    const { success, error } = await authStore.login(values.email, values.password);
    
    if (!success) {
      toast.error({title: `Error during sign in: ${error}`});
      return;
    }
    
    toast.success({title: 'Sign in successful! Redirecting to home...'});
    resetForm();
    router.push('/');
  } catch (error) {
    toast.error({title:'An unexpected error occurred. Please try again.'});
    console.error('Error during sign in:', error);
  }
});
</script>

<template>
  <div class="flex justify-center items-center h-screen bg-gray-100">
    <div class="w-full max-w-md bg-white p-8 rounded shadow">
      <h1 class="text-2xl font-bold mb-6">Sign In</h1>

      <form @submit.prevent="onSubmit">
        <div class="mb-4">
          <InputField
            name="email"              
            label="Email"
            type="email"
          />
        </div>

        <div class="mb-4">
          <InputField
            name="password"
            label="Password"
            type="Password"
          />
        </div>

        <button
          type="submit"
          class="w-full bg-red-500 text-white py-2 px-4 rounded hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
        >
          Sign In
        </button>
      </form>

      <p class="mt-4 text-center text-sm text-gray-600">
        Don't have an account? <a href="/sign-up" class="text-red-500 hover:underline">Sign Up</a>
      </p>
      <p class="mt-2 text-center text-sm text-gray-600">
        <a href="/" class="text-blue-500 hover:underline">Back to Home</a>
      </p>
    </div>
  </div>
</template>