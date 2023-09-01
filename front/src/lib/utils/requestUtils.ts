import {browser} from '$app/environment';
import type {Token, UserResponse} from '$lib/interfaces/user.interface';
import type {CustomError} from '$lib/interfaces/error.interface';
import {notificationData} from '$lib/store/notificationStore';

//import { variables } from '$lib/utils/constants';
import {formatText} from '$lib/formats/formatString';

export const browserGet = (key: string): string | undefined => {
    if (browser) {
        const item = localStorage.getItem(key);
        if (item) {
            return item;
        }
    }
    return undefined;
};

export const browserSet = (key: string, value: string): void => {
    if (browser) {
        localStorage.setItem(key, value);
    }
};


export const getCurrentUser = async (
    fetch: any,
    refreshUrl: string,
    userUrl: string
): Promise<[object, Array<CustomError>]> => {
    const jsonRes = await fetch(refreshUrl, {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            refresh: `${browserGet('refreshToken')}`
        })
    });
    const accessRefresh: Token = await jsonRes.json();
    if (accessRefresh.access) {
        const res = await fetch(userUrl, {
            headers: {
                Authorization: `Bearer ${accessRefresh.access}`
            }
        });
        if (res.status === 400) {
            const data = await res.json();
            const error = data.user.error[0];
            return [{}, error];
        }
        const response = await res.json();
        return [response.user, []];
    } else {
        return [{}, [{error: 'Refresh token is invalid...'}]];
    }
};

/*
export const handlePostRequestsWithPermissions = async (
    fetch: any,
    targetUrl: string,
    body: unknown,
    method = 'POST'
): Promise<[object, Array<CustomError>]> => {
    const res = await fetch(`${variables.BASE_API_URI}/token/refresh/`, {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            refresh: `${browserGet('refreshToken')}`
        })
    });
    const accessRefresh = await res.json();
    const jres = await fetch(targetUrl, {
        method: method,
        mode: 'cors',
        headers: {
            Authorization: `Bearer ${accessRefresh.access}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    });

    if (method === 'PATCH') {
        if (jres.status !== 200) {
            const data = await jres.json();
            console.error(`Data: ${data}`);
            const errs = data.errors;
            console.error(errs);
            return [{}, errs];
        }
        return [await jres.json(), []];
    } else if (method === 'POST') {
        if (jres.status !== 201) {
            const data = await jres.json();
            console.error(`Data: ${data}`);
            const errs = data.errors;
            console.error(errs);
            return [{}, errs];
        }
        return [jres.json(), []];
    }
};

export const UpdateField = async (
    fieldName: string,
    fieldValue: string,
    url: string
): Promise<[object, Array<CustomError>]> => {
    const userObject: UserResponse = {user: {}};
    let formData: UserResponse | any;
    if (url.includes('/user/')) {
        formData = userObject;
        formData['user'][`${fieldName}`] = fieldValue;
    } else {
        formData[`${fieldName}`] = fieldValue;
    }

    const [response, err] = await handlePostRequestsWithPermissions(fetch, url, formData, 'PATCH');
    if (err.length > 0) {
        return [{}, err];
    }
    notificationData.update(() => `${formatText(fieldName)} has been updated successfully.`);
    return [response, []];
};

 */
