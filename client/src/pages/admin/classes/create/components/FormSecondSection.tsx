// Components
import StudentTable from './StudentSelection';

export interface FormSecondSectionData {
    students: number[];
}

/**
 * A input wrapped in a FormControl. Display years as selectable options
 * starting from 2024;
 * 
 * Meant to be used under a CreateFormContext. 
 */
const SecondFormSection = (): JSX.Element => {
    return (
        <StudentTable />
    )
}


export default SecondFormSection;