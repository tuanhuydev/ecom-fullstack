<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod';
import { useForm } from 'vee-validate';
import { useRouter } from 'vue-router';
import { z } from 'zod';
import InputField from '~/components/InputField.vue';

const router = useRouter();
const toast = useToast();
const config = useRuntimeConfig();

const signUpSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  email: z.string().email('Invalid email address'),
  password: z.string().min(6, 'Password must be at least 6 characters long'),
  confirmPassword: z.string().min(6, 'Confirm Password must be at least 6 characters long')
}).refine((data) => data.password === data.confirmPassword, {
  message: "Passwords don't match",
  path: ['confirmPassword'],
});

const { handleSubmit, resetForm } = useForm({
  validationSchema: toTypedSchema(signUpSchema),
  initialValues: {
    name: '',
    email: '',
    password: '',
    confirmPassword: ''
  },
});

const onSubmit = handleSubmit(async (values) => {
  try {
    const response: Response = await $fetch(`${config.public.apiUrl}/auth/sign-up`, {
      method: 'POST',
      body: values,
    });
    if (response.ok) {
      toast.error({title: `Error during sign up`});
    } else {
      toast.success({title: 'Sign up successful! Redirecting to home...'});
      resetForm();
      router.push('/');
    }
  } catch (error) {
    toast.error({title:'An unexpected error occurred. Please try again.'});
    console.error('Error during sign up:', error);
  }
});
</script>

<template>
  <div class="flex justify-center items-center h-screen bg-gray-100">
    <div class="w-full max-w-md bg-white p-8 rounded shadow">
      <h1 class="text-2xl font-bold mb-6">Sign Up</h1>

      <form @submit.prevent="onSubmit">
        <InputField
          name="name"
          label="Name"
          type="text"
        />

        <InputField
          name="email"
          label="Email"
          type="email"
        />

        <InputField
          name="password"
          label="Password"
          type="password"
        />

        <InputField
          name="confirmPassword"
          label="Confirm Password"
          type="password"
        />

        <button
          type="submit"
          class="w-full bg-red-500 text-white py-2 px-4 rounded hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
        >
          Sign Up
        </button>
      </form>

      <p class="mt-4 text-center text-sm text-gray-600">
        Already have an account? <a href="/sign-in" class="text-red-500 hover:underline">Sign In</a>
      </p>
      <p class="mt-2 text-center text-sm text-gray-600">
        <a href="/" class="text-blue-500 hover:underline">Back to Home</a>
      </p>
    </div>
  </div>
</template>