import { Overlay, Container, Title, Text } from '@mantine/core';
import classes from './hero.module.css';

export function HeroComponent() {
	return (
		<div className={classes.hero}>
			<Overlay
				gradient="linear-gradient(180deg, rgba(0, 0, 0, 0.25) 0%, rgba(0, 0, 0, .65) 40%)"
				opacity={1}
				zIndex={0}
			/>
			<Container className={classes.container} size="md">
				<Title className={classes.title}>TrioHR</Title>
				<Text className={classes.description} size="xl" mt="xl">
					Welcome to TrioHR!
				</Text>
			</Container>
		</div>
	);
}
