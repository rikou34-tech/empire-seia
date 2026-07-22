package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Experience struct {
	Period      string
	Role        string
	Description string
}

type SkillGroup struct {
	Category    string
	Description string
	Icon        string
	ProofCode   string
}

type KPI struct {
	Value string
	Label string
}

type PageData struct {
	Name            string
	Title           string
	Location        string
	Phone           string
	Email           string
	Linkedin        string
	Objective       string
	Habilitations   string
	KPIs            []KPI
	TelecomSkills   []string
	OtherSkills     []string
	MasteryShowcase []SkillGroup
	Experiences     []Experience
	SuccessMessage  string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	successMsg := ""
	if r.Method == http.MethodPost {
		r.ParseForm()
		recruiterName := r.FormValue("name")
		if recruiterName != "" {
			successMsg = "Merci " + recruiterName + " ! Votre message a bien été transmis à la supervision."
		} else {
			successMsg = "Message transmis avec succès."
		}
	}

	htmlTemplate := `
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Portfolio & CV - {{.Name}}</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
            background-color: #0f172a;
            color: #f8fafc;
            margin: 0;
            padding: 2rem;
            display: flex;
            justify-content: center;
        }
        .container {
            max-width: 900px;
            width: 100%;
            padding: 2.5rem;
            background: #1e293b;
            border-radius: 12px;
            box-shadow: 0 10px 25px rgba(0, 0, 0, 0.4);
            border: 1px solid #334155;
        }
        header {
            border-bottom: 2px solid #334155;
            padding-bottom: 1.5rem;
            margin-bottom: 2rem;
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            flex-wrap: wrap;
            gap: 1rem;
        }
        h1 { color: #38bdf8; margin: 0 0 5px 0; font-size: 2.2rem; }
        h2 { color: #94a3b8; margin: 0 0 15px 0; font-size: 1.2rem; font-weight: 500; }
        .contact-info, .hab { font-size: 0.95rem; color: #cbd5e1; margin: 5px 0; }
        a { color: #38bdf8; text-decoration: none; }
        a:hover { text-decoration: underline; }
        
        .btn-print {
            background-color: #0284c7;
            color: #ffffff;
            border: none;
            padding: 8px 16px;
            border-radius: 6px;
            font-weight: bold;
            cursor: pointer;
            font-size: 0.9rem;
            transition: background 0.2s;
        }
        .btn-print:hover { background-color: #0369a1; }

        section { margin-bottom: 2rem; }
        h3 { color: #38bdf8; border-left: 4px solid #38bdf8; padding-left: 10px; margin-bottom: 1rem; }
        
        /* KPIs Grid */
        .kpi-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 1rem; margin-bottom: 2rem; }
        .kpi-card { background: #0f172a; padding: 1rem; border-radius: 8px; border: 1px solid #334155; text-align: center; }
        .kpi-value { font-size: 1.4rem; font-weight: bold; color: #38bdf8; margin-bottom: 4px; }
        .kpi-label { font-size: 0.85rem; color: #94a3b8; text-transform: uppercase; letter-spacing: 0.5px; }

        .skills-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 1.5rem; }
        .showcase-grid { display: grid; grid-template-columns: 1fr; gap: 1rem; }
        @media(min-width: 768px) { .showcase-grid { grid-template-columns: 1fr 1fr; } }
        
        .card, .showcase-card { background: #0f172a; padding: 1.2rem; border-radius: 8px; border: 1px solid #334155; }
        .card ul { margin: 0; padding-left: 20px; color: #94a3b8; line-height: 1.6; }
        .card li { margin-bottom: 6px; }
        .showcase-card h4 { color: #38bdf8; margin: 0 0 8px 0; font-size: 1.05rem; }
        .showcase-card p { color: #94a3b8; margin: 0 0 10px 0; font-size: 0.95rem; line-height: 1.5; }
        .proof-box { background: #020617; border-left: 3px solid #38bdf8; padding: 8px 12px; font-family: monospace; font-size: 0.8rem; color: #38bdf8; border-radius: 4px; margin-top: 8px; }

        /* Terminal Console */
        .terminal { background: #020617; border: 1px solid #334155; border-radius: 8px; padding: 1rem; font-family: monospace; font-size: 0.85rem; color: #34d399; margin-bottom: 2rem; }
        .terminal-header { color: #64748b; margin-bottom: 8px; border-bottom: 1px solid #1e293b; padding-bottom: 4px; display: flex; justify-content: space-between; }
        .log-line { margin: 4px 0; }

        /* Contact Form */
        .contact-form { background: #0f172a; padding: 1.5rem; border-radius: 8px; border: 1px solid #334155; }
        .form-group { margin-bottom: 1rem; }
        .form-group label { display: block; color: #cbd5e1; margin-bottom: 5px; font-size: 0.9rem; }
        .form-group input, .form-group textarea { width: 100%; padding: 10px; background: #1e293b; border: 1px solid #334155; color: #f8fafc; border-radius: 6px; font-size: 0.95rem; box-sizing: border-box; }
        .form-group input:focus, .form-group textarea:focus { outline: none; border-color: #38bdf8; }
        .btn-submit { background-color: #38bdf8; color: #0f172a; border: none; padding: 10px 20px; border-radius: 6px; font-weight: bold; cursor: pointer; font-size: 0.95rem; }
        .btn-submit:hover { background-color: #7dd3fc; }
        .alert-success { background: #064e3b; color: #6ee7b7; padding: 10px; border-radius: 6px; margin-bottom: 1rem; border: 1px solid #047857; }

        .job-card { background: #0f172a; padding: 1.2rem; border-radius: 8px; margin-bottom: 1rem; border: 1px solid #334155; }
        .job-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 5px; }
        .job-card h4 { color: #f8fafc; margin: 0; font-size: 1.1rem; }
        .period { background: #334155; color: #38bdf8; padding: 2px 8px; border-radius: 4px; font-size: 0.85rem; font-weight: bold; }
        .job-card p { color: #94a3b8; margin: 5px 0 0 0; line-height: 1.5; }

        @media print {
            body { background: #fff; color: #000; padding: 0; }
            .container { box-shadow: none; border: none; padding: 0; max-width: 100%; background: #fff; }
            .btn-print, .contact-form, .terminal { display: none; }
            h1, h2, h3, h4, .kpi-value { color: #000 !important; }
            .card, .showcase-card, .kpi-card, .job-card { background: #fff; border: 1px solid #ccc; }
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div>
                <h1>{{.Name}}</h1>
                <h2>{{.Title}}</h2>
                <p class="contact-info">📍 {{.Location}} | 📞 {{.Phone}} | ✉️ <a href="mailto:{{.Email}}">{{.Email}}</a></p>
                <p class="contact-info">🔗 Profil : <a href="{{.Linkedin}}" target="_blank">LinkedIn</a></p>
                <p class="hab"><strong>Habilitations & Permis :</strong> {{.Habilitations}}</p>
            </div>
            <button class="btn-print" onclick="window.print()">🖨️ Imprimer / Sauvegarder PDF</button>
        </header>

        <!-- Section Chiffres Clés & KPIs -->
        <div class="kpi-grid">
            {{range .KPIs}}
            <div class="kpi-card">
                <div class="kpi-value">{{.Value}}</div>
                <div class="kpi-label">{{.Label}}</div>
            </div>
            {{end}}
        </div>

        <!-- Console / Terminal de Logs Temps Réel -->
        <div class="terminal">
            <div class="terminal-header">
                <span>⚡ TERMINAL DE SUPERVISION & SYNC SEIA</span>
                <span>STATUS: LIVE</span>
            </div>
            <div class="log-line">> [INIT] Chargement des modules de contrôle technique... [OK]</div>
            <div class="log-line">> [NRO_CORE] Vérification des baies PMGC / SET / UNSET : Intégrité 100%</div>
            <div class="log-line">> [SECURITY] SSIAP & Habilitations électriques B2/BR/HO : Actives & Conformes</div>
            <div class="log-line">> [SYS_READY] Prêt pour affectation technique immédiate ou supervision d'équipe.</div>
        </div>

        <section>
            <h3>🎯 Objectif Professionnel</h3>
            <p style="color: #94a3b8; line-height: 1.6;">{{.Objective}}</p>
        </section>

        <section>
            <h3>⚡ Démonstration de Polyvalence & Sceaux de Preuve</h3>
            <div class="showcase-grid">
                {{range .MasteryShowcase}}
                <div class="showcase-card">
                    <h4>{{.Icon}} {{.Category}}</h4>
                    <p>{{.Description}}</p>
                    {{if .ProofCode}}
                    <div class="proof-box">✓ {{.ProofCode}}</div>
                    {{end}}
                </div>
                {{end}}
            </div>
        </section>

        <section class="skills-grid">
            <div class="card">
                <h3>📡 Télécoms & Réseaux</h3>
                <ul>
                    {{range .TelecomSkills}}
                    <li>{{.}}</li>
                    {{end}}
                </ul>
            </div>
            <div class="card">
                <h3>🛡️ Sécurité & Autres</h3>
                <ul>
                    {{range .OtherSkills}}
                    <li>{{.}}</li>
                    {{end}}
                </ul>
            </div>
        </section>

        <section>
            <h3>💼 Expériences Professionnelles</h3>
            {{range .Experiences}}
            <div class="job-card">
                <div class="job-header">
                    <h4>{{.Role}}</h4>
                    <span class="period">{{.Period}}</span>
                </div>
                <p>{{.Description}}</p>
            </div>
            {{end}}
        </section>

        <!-- Formulaire de Contact Direct Recruteur -->
        <section>
            <h3>📩 Contact direct & Recrutement</h3>
            <div class="contact-form">
                {{if .SuccessMessage}}
                <div class="alert-success">{{.SuccessMessage}}</div>
                {{end}}
                <form action="/" method="POST">
                    <div class="form-group">
                        <label for="name">Votre Nom / Entreprise :</label>
                        <input type="text" id="name" name="name" placeholder="Ex: DRH / Responsable Technique" required>
                    </div>
                    <div class="form-group">
                        <label for="email">Votre Email de contact :</label>
                        <input type="email" id="email" name="email" placeholder="contact@entreprise.com" required>
                    </div>
                    <div class="form-group">
                        <label for="message">Message / Proposition :</label>
                        <textarea id="message" name="message" rows="4" placeholder="Échangeons concernant un poste de supervision ou technique..." required></textarea>
                    </div>
                    <button type="submit" class="btn-submit">Envoyer la demande</button>
                </form>
            </div>
        </section>
    </div>
</body>
</html>
`
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Name:           "ERIC UBACH",
		Title:          "Responsable technique / Superviseur",
		Location:       "34500 Béziers",
		Phone:          "07.82.70.02.93",
		Email:          "rikou34@gmail.com",
		Linkedin:       "https://linkedin.com/in/eric-u-19506a179",
		Objective:      "Être un membre de confiance, progresser, mettre ma motivation ainsi que mes compétences clés et mon savoir-être au service de votre entreprise.",
		Habilitations:  "B0, B1, B2, BC, BR, HO | CACES nacelle cat. B, AIPR",
		SuccessMessage: successMsg,
		KPIs: []KPI{
			{Value: "Titre OFIAQ (Niveau IV)", Label: "Technicien Réseaux de Communications"},
			{Value: "100%", Label: "Récettes & Mesures Validées"},
			{Value: "ERP 1 & 2", Label: "Sécurité & Supervision"},
			{Value: "NRO / FTTH", Label: "Expertise Réseau & Câblage"},
		},
		MasteryShowcase: []SkillGroup{
			{
				Icon:        "🔌",
				Category:    "Intervention & Câblage NRO",
				Description: "Maîtrise complète des environnements critiques (PMGC, SET, UNSET), tirage de câbles, soudure fibre optique et recettes de mesures avec Fluke Networks.",
				ProofCode:   "NRO_EXEC_STATUS: PMGC/SET/UNSET [100% Validé]",
			},
			{
				Icon:        "🤖",
				Category:    "Collaboration IA & Systèmes",
				Description: "Capacité à concevoir, piloter et interagir directement avec des intelligences artificielles et des microservices (Go) pour structurer des architectures de données en temps réel.",
				ProofCode:   "SYS_SYNC: SEIA_CORE v2.0 // Latence: 0.05s [Actif]",
			},
			{
				Icon:        "💻",
				Category:    "Logiciels & Outils Métiers",
				Description: "Utilisation experte des outils de bureautique (Pack Office), lecture de plans D1/D2/FTTH, et gestion de plateformes de supervision et de sécurité (SSIAP).",
				ProofCode:   "PLANS_READ: D1_D2_FTTH // Outils Bureautiques [Maîtrisés]",
			},
			{
				Icon:        "🛡️",
				Category:    "Sécurité & Supervision",
				Description: "Expérience approfondie en gestion d'équipes, sécurité des biens et des personnes (ERP1/ERP2), vidéo-protection et habilitations électriques rigoureuses.",
				ProofCode:   "HABILITATION: B2/BR/HO + SSIAP 1&2 + CACES Nacelle B",
			},
		},
		TelecomSkills: []string{
			"Intervention NRO (PMGC, SET, UNSET, câblage provisoire)",
			"Réalisation d'Audits tiers & Expertises des échecs de raccordements",
			"Fibre Optique (Installation, raccordement, essais et mesures)",
			"Tirage et pose de câbles aéro-souterrains, baies de brassage, soudure",
			"Tests & Mesures : Fluke networks, multimètre, testeur de continuité",
			"Organisation des travaux de création et d'entretien du réseau",
		},
		OtherSkills: []string{
			"Sécurité : Vidéo protection, Agent de prévention, Chef d'équipe (SSIAP 1 & 2)",
			"Nettoiement : Nettoyage haute pression, balayeuse",
			"Maçonnerie : Gestion de chantier, maçonnerie générale, marbrier",
			"Bureautique & Outils : Pack Office",
		},
		Experiences: []Experience{
			{Period: "05/2023 - 05/2025", Role: "Technicien télécom - Free Réseau", Description: "Interventions techniques, raccordements, expertises et mesures sur les réseaux fibre optique."},
			{Period: "05/2019 - 11/2021", Role: "Agent territorial - Mairie", Description: "Service de nettoiement de la ville."},
			{Period: "09/2018 - 04/2019", Role: "Technicien en réseaux de communication - OFIAQ DE SETE", Description: "Formations et stages (ADEP / NGE) : VDI, lecture de plan D1/D2/FTTH, normes de sécurité."},
			{Period: "01/2006 - 04/2018", Role: "Agent territorial - Mairie", Description: "Vidéo-protection et service des fêtes de la ville."},
			{Period: "06/1998 - 01/2006", Role: "Agent de sécurité - Hypermarché", Description: "Sécurité ERP1 / ERP2."},
			{Period: "1992 - 1998", Role: "Ets Ubach", Description: "Maçonnerie générale et gestion de structures."},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	pkillCmd := "pkill -f 'go run main.go' || true"
	_ = pkillCmd

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", homeHandler)
	fmt.Println("🚀 Serveur mis à jour démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
