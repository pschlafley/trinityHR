import { useLocation } from 'react-router-dom';
import { Container, Group, Anchor } from '@mantine/core';
import classes from './navbar.module.css';
// import logo from '../../assets/TrinityLogo.png';

export function NavBar() {
	const windowPath = useLocation().pathname;

	const homePath = windowPath == '/' ? true : false;
	const accountsPath = windowPath == '/accounts' ? true : false;
	const loginPath = windowPath == '/login' ? true : false;

	return (
		<header className={classes.header}>
			<Container size="md" className={classes.inner}>
				<Group gap={5} visibleFrom="xs">
					<Anchor
						key="home"
						href="/"
						className={
							homePath
								? `${classes.active} ${classes.linkActive}`
								: `${classes.link}`
						}
					>
						Home
					</Anchor>
					<Anchor
						key="accounts"
						href="/accounts"
						className={
							accountsPath
								? `${classes.active} ${classes.linkActive}`
								: `${classes.link}`
						}
					>
						Accounts
					</Anchor>
					<Anchor
						key="login"
						href="/login"
						className={
							loginPath
								? `${classes.active} ${classes.linkActive}`
								: `${classes.link}`
						}
					>
						Login
					</Anchor>
				</Group>
			</Container>
		</header>
	);
}
