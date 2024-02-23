import { z } from 'zod';
import { fail } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';

const loginSchema = z.object({
	email: z.string().email(),
	password: z.string().min(8)
});

const registerSchema = z.object({
	email: z.string().email(),
	password: z.string.min(8),
	confirmPassword: z.string.min(8)
});

export const load = async () => {
	// Different schemas, no id required.
	const loginForm = await superValidate(zod(loginSchema));
	const registerForm = await superValidate(zod(registerSchema));

	// Return them both
	return { loginForm, registerForm };
};

export const actions = {
	login: async ({ request }) => {
		const loginForm = await superValidate(request, zod(loginSchema));

		if (!loginForm.valid) return fail(400, { loginForm });

		// TODO: Login user
		return message(loginForm, 'Login form submitted');
	},

	register: async ({ request }) => {
		const registerForm = await superValidate(request, zod(registerSchema));

		if (!registerForm.valid) return fail(400, { registerForm });

		// TODO: Register user
		return message(registerForm, 'Register form submitted');
	}
};