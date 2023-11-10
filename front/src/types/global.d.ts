declare module 'svelte-cookie' {
    export function setCookie(
        name: string,
        value: string,
        options?: import('cookie').CookieSerializeOptions
    ): void;

    export function getCookie(name: string): string | undefined;

    export function deleteCookie(name: string): void;
}