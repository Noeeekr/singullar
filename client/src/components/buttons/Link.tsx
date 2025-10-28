// Components
import { Link } from 'react-router-dom'
import Button from './Default';

// Features
import { useLocation} from "react-router-dom"
import { useAppDispatch } from '../../slices/store';
import { incrementUrlVisitedCount } from '../../slices/userSlice'

// Types
import type { ButtonsProps } from "@components/buttons/Default"

export interface LinkButtonProps extends ButtonsProps {
    href: string
    type?: "link"
}

const LinkButton = ({ href, variant = "button", ...props}: LinkButtonProps): JSX.Element => {
    const dispatch = useAppDispatch();
    const location = useLocation();
    if (props.effects == undefined) props.effects = { enableSelectEffect: location.pathname == href };
    if (props.effects?.enableSelectEffect == undefined) props.effects.enableSelectEffect = location.pathname == href;
    
    return (
        <Link 
            to={href}
            style={{ textDecoration: 'none' }}
            onClick={() => (dispatch(incrementUrlVisitedCount(href)))}
        >
            <Button {...props} variant={variant} />
        </Link>
    )
}

export default LinkButton;