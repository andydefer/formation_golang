// Package paiement provides a flexible payment processing system with support
// for multiple payment methods including credit cards, PayPal, and cryptocurrency.
package paiement

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Logger pour le package paiement
var (
	paymentInfoLog  *log.Logger
	paymentErrorLog *log.Logger
	initOnce        sync.Once
)

func init() {
	initOnce.Do(func() {
		// Ouvre le fichier de log en mode append
		fichier, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Impossible d'ouvrir le fichier de log:", err)
		}

		paymentInfoLog = log.New(fichier, "PAYMENT-INFO: ", log.Ldate|log.Ltime)
		paymentErrorLog = log.New(fichier, "PAYMENT-ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

// Payeur interface définissant le contrat pour tout mode de paiement
type Payeur interface {
	Payer(montant float64) error
	GetNom() string
}

// Creditable interface pour les moyens de paiement qui supportent le crédit
type Creditable interface {
	Credit(montant float64) error
}

// CarteBancaire structure
type CarteBancaire struct {
	Numero    string
	DateExp   string
	Titulaire string
	Solde     float64
	mu        sync.RWMutex // Pour la sécurité concurrente
}

// NouvelleCarteBancaire crée une nouvelle carte avec solde initial
func NouvelleCarteBancaire(numero, dateExp, titulaire string) *CarteBancaire {
	paymentInfoLog.Printf("Nouvelle carte créée pour %s (solde: 1000.00€)", titulaire)
	return &CarteBancaire{
		Numero:    numero,
		DateExp:   dateExp,
		Titulaire: titulaire,
		Solde:     1000.00,
	}
}

func (c *CarteBancaire) Payer(montant float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if montant <= 0 {
		err := fmt.Errorf("montant invalide: %.2f€", montant)
		paymentErrorLog.Printf("Carte %s - %v", c.Titulaire, err)
		return err
	}
	if c.Solde < montant {
		err := fmt.Errorf("solde insuffisant: %.2f€ disponible, %.2f€ requis", c.Solde, montant)
		paymentErrorLog.Printf("Carte %s - %v", c.Titulaire, err)
		return err
	}
	c.Solde -= montant
	paymentInfoLog.Printf("Carte %s - Paiement de %.2f€ effectué. Nouveau solde: %.2f€",
		c.Titulaire, montant, c.Solde)
	return nil
}

func (c *CarteBancaire) GetNom() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Titulaire
}

// Credit implémente l'interface Creditable pour la carte bancaire
func (c *CarteBancaire) Credit(montant float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if montant <= 0 {
		err := fmt.Errorf("montant de crédit invalide: %.2f€", montant)
		paymentErrorLog.Printf("Carte %s - %v", c.Titulaire, err)
		return err
	}
	c.Solde += montant
	paymentInfoLog.Printf("Carte %s - Crédit de %.2f€ effectué. Nouveau solde: %.2f€",
		c.Titulaire, montant, c.Solde)
	return nil
}

// GetNumeroMasque retourne le numéro masqué pour l'affichage
func (c *CarteBancaire) GetNumeroMasque() string {
	if len(c.Numero) >= 4 {
		return "****" + c.Numero[len(c.Numero)-4:]
	}
	return "****"
}

// PayPal structure
type PayPal struct {
	Email string
	Solde float64
	mu    sync.RWMutex
}

// NouveauPayPal crée un nouveau compte PayPal avec solde initial
func NouveauPayPal(email string) *PayPal {
	paymentInfoLog.Printf("Nouveau compte PayPal créé pour %s (solde: 500.00€)", email)
	return &PayPal{
		Email: email,
		Solde: 500.00,
	}
}

func (p *PayPal) Payer(montant float64) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if montant <= 0 {
		err := fmt.Errorf("montant invalide: %.2f€", montant)
		paymentErrorLog.Printf("PayPal %s - %v", p.Email, err)
		return err
	}
	if p.Solde < montant {
		err := fmt.Errorf("solde insuffisant: %.2f€ disponible, %.2f€ requis", p.Solde, montant)
		paymentErrorLog.Printf("PayPal %s - %v", p.Email, err)
		return err
	}
	p.Solde -= montant
	paymentInfoLog.Printf("PayPal %s - Paiement de %.2f€ effectué. Nouveau solde: %.2f€",
		p.Email, montant, p.Solde)
	return nil
}

func (p *PayPal) GetNom() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.Email
}

// Credit implémente l'interface Creditable pour PayPal
func (p *PayPal) Credit(montant float64) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if montant <= 0 {
		err := fmt.Errorf("montant de crédit invalide: %.2f€", montant)
		paymentErrorLog.Printf("PayPal %s - %v", p.Email, err)
		return err
	}
	p.Solde += montant
	paymentInfoLog.Printf("PayPal %s - Crédit de %.2f€ effectué. Nouveau solde: %.2f€",
		p.Email, montant, p.Solde)
	return nil
}

// Crypto structure
type Crypto struct {
	Adresse string
	Solde   float64
	mu      sync.RWMutex
}

// NouvelleCrypto crée un nouveau portefeuille crypto avec solde initial
func NouvelleCrypto(adresse string) *Crypto {
	paymentInfoLog.Printf("Nouveau portefeuille Crypto créé pour %s (solde: 1000.00€)", adresse)
	return &Crypto{
		Adresse: adresse,
		Solde:   1000.00,
	}
}

func (c *Crypto) Payer(montant float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if montant <= 0 {
		err := fmt.Errorf("montant invalide: %.2f€", montant)
		paymentErrorLog.Printf("Crypto %s - %v", c.Adresse, err)
		return err
	}
	frais := montant * 0.02
	total := montant + frais
	if c.Solde < total {
		err := fmt.Errorf("solde insuffisant: %.2f€ disponible, %.2f€ requis (dont %.2f€ de frais)",
			c.Solde, total, frais)
		paymentErrorLog.Printf("Crypto %s - %v", c.Adresse, err)
		return err
	}
	c.Solde -= total
	paymentInfoLog.Printf("Crypto %s - Paiement de %.2f€ + frais %.2f€ = %.2f€. Nouveau solde: %.2f€",
		c.Adresse, montant, frais, total, c.Solde)
	return nil
}

func (c *Crypto) GetNom() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Adresse
}

// Credit implémente l'interface Creditable pour Crypto
func (c *Crypto) Credit(montant float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if montant <= 0 {
		err := fmt.Errorf("montant de crédit invalide: %.2f€", montant)
		paymentErrorLog.Printf("Crypto %s - %v", c.Adresse, err)
		return err
	}
	c.Solde += montant
	paymentInfoLog.Printf("Crypto %s - Crédit de %.2f€ effectué. Nouveau solde: %.2f€",
		c.Adresse, montant, c.Solde)
	return nil
}

// GetFrais retourne le montant des frais pour un paiement donné
func (c *Crypto) GetFrais(montant float64) float64 {
	return montant * 0.02
}

// EffectuerPaiement fonction générique de paiement avec logging
func EffectuerPaiement(p Payeur, montant float64) {
	fmt.Printf("\n--- Paiement de %.2f€ avec %s ---\n", montant, p.GetNom())

	err := p.Payer(montant)
	if err != nil {
		fmt.Printf("❌ Échec du paiement: %v\n", err)
		return
	}

	fmt.Printf("✅ Paiement de %.2f€ effectué avec succès\n", montant)

	// Affichage spécifique selon le type
	switch v := p.(type) {
	case *CarteBancaire:
		fmt.Printf("   Nouveau solde: %.2f€\n", v.Solde)
	case *PayPal:
		fmt.Printf("   Nouveau solde: %.2f€\n", v.Solde)
	case *Crypto:
		frais := v.GetFrais(montant)
		fmt.Printf("   Frais: %.2f€, Total: %.2f€, Nouveau solde: %.2f€\n", frais, montant+frais, v.Solde)
	}
}

// AfficherInfos affiche les détails spécifiques à chaque mode de paiement
func AfficherInfos(p interface{}) {
	switch v := p.(type) {
	case *CarteBancaire:
		fmt.Printf("[CB] Titulaire: %s, Numéro: %s, Date exp: %s, Solde: %.2f€\n",
			v.Titulaire, v.GetNumeroMasque(), v.DateExp, v.Solde)
	case *PayPal:
		fmt.Printf("[PayPal] Email: %s, Solde: %.2f€\n", v.Email, v.Solde)
	case *Crypto:
		fmt.Printf("[Crypto] Adresse: %s, Solde: %.2f€\n", v.Adresse, v.Solde)
	default:
		fmt.Printf("Type inconnu: %T\n", v)
	}
}
