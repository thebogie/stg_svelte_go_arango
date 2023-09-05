import {sveltekit} from '@sveltejs/kit/vite';
import type {UserConfig} from 'vitest/config';

const config: UserConfig = {
    server: {
        port: 50003,
        host: '0.0.0.0',

    },
    plugins: [sveltekit()],
    test: {
        include: ['src/**/*.{test,spec}.{js,ts}']
    }
};

export default config;
