// Components
import Typography from '@mui/material/Typography';
import InstitutionSelection from './InstitutionSelection';

const PageInstitutionSelection = (): JSX.Element => {
    return(
        <div>
            <Typography marginBottom={2} color="primary.purpleLight" variant="h5" fontWeight="600">
                Busque uma escola
            </Typography>
            <InstitutionSelection />
        </div>
    )
}

export default PageInstitutionSelection;