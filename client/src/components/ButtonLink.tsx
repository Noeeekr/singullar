// Components
import { Link } from 'react-router-dom'
import Button from './ButtonSolid'
import Buttons from './Button'

// Features

import { useAppDispatch } from '../slices/store';
import { incrementUrlVisitedCount } from '../slices/userSlice'

// Types
import type { ISideMenuButtonProps, IButtonBaseProps } from './Button'
import type { ISideMenuLinkButton } from '../models/buttonProps'

export interface ISideMenuLinkProps extends ISideMenuLinkButton, IButtonBaseProps {
    href: string,
    children?: string,
};

const LinkButton = (props: ISideMenuLinkProps): JSX.Element => {
    const { href, children, variant } = props;
    
    const dispatch = useAppDispatch();
    
    const clonedProps: ISideMenuButtonProps = { ...props, type: "button" }

    return (
        <Link 
            to={href}
            style={{ textDecoration: 'none' }}
            onClick={() => (dispatch(incrementUrlVisitedCount(href)))}
        >
            {
                variant === undefined
                ? <Button {...clonedProps}>
                    { children ? children : "Children not found"}
                </Button>
                : <Buttons {...clonedProps}/>
            }
        </Link>
    )
}

export default LinkButton;