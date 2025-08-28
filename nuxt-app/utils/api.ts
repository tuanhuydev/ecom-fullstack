import { useAuthStore } from '~/stores/authStore';

interface FetchOptions extends RequestInit {
  body?: BodyInit | null;
}

/**
 * Enhanced fetch utility that automatically includes auth token
 * and handles common API request configurations
 */
export async function apiFetch<T = unknown>(url: string, options: FetchOptions = {}): Promise<T> {
  const authStore = useAuthStore();
  const config = useRuntimeConfig();
  
  // Create headers with authorization if token exists
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
    ...options.headers,
  };
  
  // Add auth token if available
  if (authStore.token) {
    headers['Authorization'] = `Bearer ${authStore.token}`;
  }

  // Prepare the request config
  const requestConfig: FetchOptions = {
    ...options,
    headers,
  };

  // Stringify the body if it's an object
  if (requestConfig.body && typeof requestConfig.body === 'object') {
    requestConfig.body = JSON.stringify(requestConfig.body);
  }

  try {
    const response = await fetch(`${config.public.apiUrl}${url}`, requestConfig);
    
    // Check if the response is ok
    if (!response.ok) {
      // Handle 401 Unauthorized - token might be expired
      if (response.status === 401) {
        authStore.logout();
        navigateTo('/sign-in');
      }
      
      // Try to parse error response
      const errorData = await response.json().catch(() => null);
      throw new Error(
        errorData?.error || 
        errorData?.message || 
        `Request failed with status ${response.status}`
      );
    }
    
    // Parse JSON response if unknown
    const contentType = response.headers.get('content-type');
    if (contentType && contentType.includes('application/json')) {
      return await response.json();
    }
    
    return await response.text() as unknown as T;
  } catch (error: unknown) {
    console.error(`API fetch error for ${url}:`, error);
    throw error;
  }
}