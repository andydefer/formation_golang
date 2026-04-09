.PHONY: generate-ai-diff
generate-ai-diff:
	@echo "🤖 Génération d'un diff pour l'analyse IA..."
	@echo ""
	@echo "📁 Quels dossiers souhaitez-vous analyser ?"
	@echo "   Exemples: app/Http/Controllers app/Models app/Services resources/js"
	@echo "   (laissez vide pour tous les dossiers modifiés)"
	@printf "Dossiers: "
	@read folders; \
	echo ""; \
	echo "📋 Récupération des fichiers modifiés dans le dernier commit..."; \
	if [ -z "$$folders" ]; then \
		files=$$(git diff --name-only HEAD~1..HEAD | grep -E '\.(php|ts|tsx|js|jsx)$$' | tr '\n' ' '); \
	else \
		files=""; \
		for folder in $$folders; do \
			new_files=$$(git diff --name-only HEAD~1..HEAD | grep "^$$folder" | grep -E '\.(php|ts|tsx|js|jsx)$$' | tr '\n' ' '); \
			files="$$files $$new_files"; \
		done; \
	fi; \
	if [ -z "$$files" ]; then \
		echo "❌ Aucun fichier modifié trouvé dans les dossiers sélectionnés."; \
		exit 1; \
	fi; \
	echo "✅ Fichiers trouves: $$files"; \
	echo ""; \
	echo "📝 Generation du diff..."; \
	output_file="ai-diff-$$(date +%Y%m%d_%H%M%S).txt"; \
	{ \
		echo "Tu es un expert en revue de code et en conventions de commits (Conventional Commits)."; \
		echo ""; \
		echo "A partir du diff Git ci-dessous, fais les choses suivantes :"; \
		echo ""; \
		echo "1. Propose un nom de commit clair et concis en anglais"; \
		echo "   avec le format <type>(<scope>): <description>,"; \
		echo "   en respectant les Conventional Commits"; \
		echo "   (ex: feat:, fix:, refactor:, test:, chore:, docs:)."; \
		echo ""; \
		echo "2. Redige un resume du travail effectue en quelques phrases,"; \
		echo "   oriente metier et technique."; \
		echo ""; \
		echo "3. Donne une liste d'exemples concrets de changements, en t'appuyant sur le diff :"; \
		echo "   - methodes ajoutees, modifiees ou supprimees"; \
		echo "   - responsabilites deplacees ou clarifiees"; \
		echo "   - ameliorations de validation, de logique ou de structure"; \
		echo "   - impacts fonctionnels eventuels"; \
		echo ""; \
		echo "Contraintes :"; \
		echo "   - Ne decris que ce qui est reellement visible dans le diff"; \
		echo "   - Sois precis, factuel et structure"; \
		echo "   - Evite les suppositions"; \
		echo "   - Utilise un ton professionnel"; \
		echo ""; \
		echo "4. SI et SEULEMENT SI les changements sont cassants (breaking changes) :"; \
		echo "   - Genere une entree de CHANGELOG conforme a Keep a Changelog et SemVer."; \
		echo "   - Le changelog doit apparaître APRES les recommandations ci-dessus."; \
		echo "   - Utilise STRICTEMENT la structure suivante :"; \
		echo ""; \
		echo "     ## [X.0.0] - YYYY-MM-DD"; \
		echo "     ### Changed"; \
		echo "     - Description claire du changement cassant"; \
		echo ""; \
		echo "     ### Removed (si applicable)"; \
		echo "     - API, methode ou comportement supprime"; \
		echo ""; \
		echo "     ### Security (si applicable)"; \
		echo "     - Impact securite lie au changement"; \
		echo ""; \
		echo "   - Ne genere PAS de changelog si aucun breaking change n'est detecte."; \
		echo "   - N'invente PAS de version."; \
		echo ""; \
		echo "Voici le diff :"; \
		echo ""; \
		git diff HEAD~1..HEAD -- $$files; \
	} > $$output_file; \
	echo ""; \
	echo "✅ Diff genere dans: $$output_file"; \
	echo "📊 Taille: $$(wc -l < $$output_file) lignes"; \
	echo ""; \
	echo "🚀 Tu peux maintenant copier le contenu du fichier et le coller dans ton IA preferee (ChatGPT, Claude, etc.)"
