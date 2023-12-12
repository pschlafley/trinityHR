import { Title, Text, Overlay } from '@mantine/core';
import classes from './NotAuthenticate.module.css';

export const NotAuthenticated = () => {
	return (
		<div className={classes.wrapper}>
			<Overlay color="#000" opacity={0.65} zIndex={1} />

			<div className={classes.inner}>
				<Title className={classes.title}>
					<Text>
						You are not authorized to view this content, please login or signup!
					</Text>
				</Title>
			</div>
		</div>
	);
};
