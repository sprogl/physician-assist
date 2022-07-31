import {AppBar, Button, IconButton, Toolbar, Typography} from "@material-ui/core"
import MenuIcon from "@material-ui/icons/Menu"
import {makeStyles} from "@material-ui/core/styles"
import { Link } from 'react-router-dom'

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
    },
}))

export default function TopAppBar() {
    const classes = useStyles()

    return (
        <div className={classes.root}>
            <AppBar position="static">
                <Toolbar>
                    <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" className={classes.title}>
                        DocAssist
                    </Typography>
                    <Link to="/login" color="inherit">Login</Link>
                </Toolbar>
            </AppBar>
        </div>
    )
}