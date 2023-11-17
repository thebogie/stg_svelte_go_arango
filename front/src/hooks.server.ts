import type { Handle } from '@sveltejs/kit';

const handle: Handle = async ({ event, resolve }) => {
	const authCookie = event.cookies.get('Authorization');
	console.log('Handling Event' + JSON.stringify(event));
	const loggedinplayer = event.cookies.get('loggedinplayer');


	if (loggedinplayer) {
		// Remove prefix
		//const token = authCookie.split(' ')[1];

		event.locals.player = loggedinplayer;



		try {


			/*
            const jwtUser = jwt.verify(token, JWT_ACCESS_SECRET);
            if (typeof jwtUser === 'string') {
                throw new Error('Something went wrong');
            }

            const user = await db.user.findUnique({
                where: {
                    id: jwtUser.id
                }
            });

            if (!user) {
                throw new Error('User not found');
            }

            const sessionUser = {
                id: user.id,
                email: user.email
            };

            event.locals.player = sessionUser;


             */
		} catch (error) {
			console.error(error);
		}
	}

	return await resolve(event);
};

export {handle};