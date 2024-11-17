// Components
import { Link } from 'react-router-dom'
import Button from './Button'

// Features

import { useAppDispatch } from '../slices/store';
import { incrementUrlVisitedCount } from '../slices/userSlice'

// Types
import type { ISideMenuButtonProps, IButtonBaseProps } from './Button'
import type { ISideMenuLinkButton } from '../types/buttonProps'

export interface ISideMenuLinkProps extends ISideMenuLinkButton, IButtonBaseProps {
    href: string,
};

const LinkButton = (props: ISideMenuLinkProps): JSX.Element => {
    const { href } = props;
    
    const dispatch = useAppDispatch();
    
    const clonedProps: ISideMenuButtonProps = { ...props, type: "button" }

    return (
        <Link 
            to={href}
            style={{ textDecoration: 'none' }}
            onClick={() => (dispatch(incrementUrlVisitedCount(href)))}
        >
            <Button {...clonedProps}/>
        </Link>
    )
}

export default LinkButton