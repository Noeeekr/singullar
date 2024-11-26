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
import type { InstitutionsState } from '../../../../../types/data';

export async function Loader() {
    const institutions = await fetch("http://localhost:8000/api/user/institutions",{
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
        return res.data;
    });

    return institutions;
}

const PageInstitutionSelection = (): JSX.Element => {
    const institutions = useLoaderData() as InstitutionsState[];
    
    const dispatch = useAppDispatch();
    useEffect(() => {
        dispatch(ActionSetInstitutions(institutions))
    },[])

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

                                if (e.target.value) {
                                    navigate("/admin/classes/" + e.target.value);
                                }
                            }}
                            value={selectValue}
                        >
                            {
                                institutions.map((inst) => {
                                    return <MenuItem
                                        value={inst.id}
                                        key={inst.id + inst.name}
                                    >{inst.name}</MenuItem>
                                })
                            }
                            {
                                Boolean(institutions.length)
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