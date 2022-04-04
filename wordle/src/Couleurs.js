import React from "react";

function MotCouleur({ item }) {

    if (item.couleurs && item.mot) {
        return (<>
            <div>
                {item.mot.split("").map((c, i) => (
                    <span key={i} style={{ color: item.couleurs[c] }}>{c}</span>
                ))}
            </div>
        </>)
    }
}

export default MotCouleur;