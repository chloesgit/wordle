import React from "react"
import MotCouleur from './Couleurs.js'
function Trouver() {
    const [mot_entre, set_mot_entre] = React.useState('')
    const [liste_mot_entre, set_liste_mot_entre] = React.useState([])
    async function resultat(e) {
        e.preventDefault();
        const data = await fetch("http://localhost:8000/" + mot_entre)
        const response = await data.json()
        set_liste_mot_entre([...liste_mot_entre, { "mot": mot_entre, "couleurs": response }])
    }

    return (<div>
        WORDLE
        <input type="text" value={mot_entre} onChange={e => set_mot_entre(e.target.value)} />
        <button type="button" onClick={resultat}>Tester le mot écrit</button>
        {liste_mot_entre.map((item, index) => {
            return (<div key={index}> <MotCouleur item={item} /> </div>)

        })}
    </div>)

}






export default Trouver