import React from "react"
import MotCouleur from './Couleurs.js'
function Trouver() {
    const [mot_entre, set_mot_entre] = React.useState('')
    const [erreur, set_erreur] = React.useState('')
    const [liste_mot_entre, set_liste_mot_entre] = React.useState([])

    async function resultat(e) {
        e.preventDefault();
        const data = await fetch("http://localhost:8000/" + mot_entre)
        if (data.status == 500) {
            const response = await data.text()
            set_erreur("Le mot n'a pas la bonne longueur, il devrait être de longueur " + response)

        } else {
            set_erreur("")
            const response = await data.json()
            set_liste_mot_entre([...liste_mot_entre, { "mot": mot_entre, "couleurs": response }])
        }

    }


    return (<div>
        WORDLE
        <p>{erreur}</p>
        <input type="text" value={mot_entre} onChange={e => set_mot_entre(e.target.value)} />
        <button type="button" onClick={resultat}>Tester le mot écrit</button>
        {liste_mot_entre.map((item, index) => {
            return (<div key={index}> <MotCouleur item={item} /> </div>)

        })}
    </div>)

}






export default Trouver