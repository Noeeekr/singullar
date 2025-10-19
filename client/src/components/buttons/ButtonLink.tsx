// Components
import { Link } from 'react-router-dom'
import Button from './Button';

// Features

import { useAppDispatch } from '../../slices/store';
import { incrementUrlVisitedCount } from '../../slices/userSlice'

// Types
import type { ButtonsProps } from "@components/buttons/Button"

export interface LinkButtonProps extends ButtonsProps {
    href: string
    type: "link"
}

const LinkButton = ({ href, variant = "button", ...props}: LinkButtonProps): JSX.Element => {
    const dispatch = useAppDispatch();
    
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