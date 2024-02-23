import { getContext, hasContext, setContext } from "svelte";
import { readable, writable } from "svelte/store";

// context for any type of store
export const getSharedStore = <T, A>(
    name: string,
    fn: (value?: A) => T,
    defaultValue?: A,
) => {
    if (hasContext(name)) {
        return getContext<T>(name);
    }
    const _value = fn(defaultValue);
    setContext(name, _value);
    return _value;
};

// writable store context
export const useWritable = <T>(name: string, value: T) =>
    getSharedStore(name, writable, value);

// readable store context
export const useReadable = <T>(name: string, value: T) =>
    getSharedStore(name, readable, value);