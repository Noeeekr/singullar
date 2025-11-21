import Questions from "../../../components/Questions"

export interface SecondSectionFormState {
    question_list_question_ids: number[]
} 

export interface SecondSectionProps {
    title: string
    onParentClick: (cb: () => void) => void
    parentButtonTitle: string
}

export default function({ onParentClick }: SecondSectionProps): JSX.Element {
    return(
        <Questions onParentClick={onParentClick} />
    )
}