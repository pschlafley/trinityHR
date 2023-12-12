import { JwtPayload, jwtDecode } from 'jwt-decode';
import Cookies from 'js-cookie';

export const setToken = (loginToken: string): string | undefined | Error => {
	try {
		return Cookies.set('token', loginToken);
	} catch (error) {
		return Error('Error: could not set token');
	}
};

export const getToken = (): string => {
	const token = Cookies.get('token');

	if (!token) {
		return '';
	}
	return token;
};

export const getProfile = (): JwtPayload | undefined => {
	if (getToken() == undefined) {
		return undefined;
	} else {
		return jwtDecode<JwtPayload>(getToken());
	}
};

export const loggedIn = (): boolean => {
	const token: string = getToken();
	return !!token;
};
