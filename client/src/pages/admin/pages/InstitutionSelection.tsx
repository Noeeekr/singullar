// Features
import { useState, useEffect } from 'react';
import { useNavigate, useLoaderData } from 'react-router-dom'
import { useAppDispatch } from '@slices/store';
import { setInstitutions as ActionSetInstitutions } from '@slices/institutionsSlice';

// Components
import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';

// Types
import type { InstitutionsState } from '../../../types/data';

export async function Loader() {
    const institutions = await fetch("http://localhost:8000/api/user/",{
        "headers": {
            "Content-Type": "application/json"
        },
        "credentials": "include",
    })
    .then(res => (res.json()))
    .then(res => {
            if (res.error) {
                return []
            }         
            console.log(res.data)
            return res.data;
    });

    return institutions;
}

/**
 * A page with an input to choose a school to manage
 * 
 * @param nextUrl the href of the page where it should send after selection.
 */
const PageInstitutionSelection = ({ nextUrl }: { nextUrl: string }): JSX.Element => {
    const institutions = useLoaderData() as InstitutionsState[];

    const dispatch = useAppDispatch();
    useEffect(() => {
        dispatch(ActionSetInstitutions({ current: null, institutions: institutions }))
    },[dispatch, institutions])

    const [selectValue, setSelectValue] = useState("");

    const navigate = useNavigate();

    return (
        <div>
            <Typography marginBottom={2} color="primary.purpleLight" variant="h5" fontWeight="600">
                Busque uma escola
            </Typography>
            <div>
                <Typography component="p">
                    Você tem {institutions.length ? institutions.length : "?"} escolas em sua rede
                </Typography>
                <form>
                    <FormControl
                        aria-labelledby="school-select-label"
                        sx={{
                            marginY: 1
                        }}
                    >
                        <InputLabel id="school-select-label" sx={{
                            zIndex: 5,
                        }}>
                            Selecione uma escola
                        </InputLabel>
                        <Select
                            labelId="school-select-label"
                            id="school-select"
                            label="Selecione uma escola"
                            onChange={(e) => {
                                setSelectValue(e.target.value)
                            }}
                            value={selectValue}
                        >
                            {
                                institutions.map((inst) => {
                                    return <MenuItem
                                        value={inst.id}
                                        key={inst.id + inst.name}
                                        onClick={() => {
                                            const i = institutions.find(ins => ins.id == inst.id);
                                            
                                            dispatch(ActionSetInstitutions({ current: i ? i : null, institutions: institutions }));
                                            navigate(nextUrl);
                                        }}
                                    >{inst.name}</MenuItem>
                                })
                            }
                            {
                                institutions.length
                                    ? []
                                    : <MenuItem
                                        value=""
                                        key="sometextthatisnotainstname"
                                    >Procurando..</MenuItem>
                            }
                        </Select>
                    </FormControl>
                </form>
            </div>
        </div>
    )
}

export default PageInstitutionSelection;