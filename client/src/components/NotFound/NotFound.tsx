import { Title, Text, Button, Container, Anchor, Group } from '@mantine/core';
import classes from './NotFound.module.css';
import { useRouteError } from 'react-router-dom';

export function NotFound() {
	const error: unknown = useRouteError();
	console.error(error);

	return (
		<Container className={classes.root}>
			<div className={classes.label}>404</div>
			<Title className={classes.title}>You have found a secret place.</Title>
			<Text c="dimmed" size="lg" ta="center" className={classes.description}>
				Unfortunately, this is only a 404 page. You may have mistyped the
				address, or the page has been moved to another URL.
			</Text>
			<Group justify="center">
				<Button variant="subtle" size="md">
					<Anchor href="/">Take me back to home page</Anchor>
				</Button>
			</Group>
		</Container>
	);
}
