# ===================================================
# Generic Makefile (PHP/Laravel specific tools removed)
# ===================================================

# ---------------------------------------------------
# Source Configuration
# ---------------------------------------------------
SOURCE_DIRS = src config database tests
IGNORED_FILES = CHANGED_FILES.md FILES_CHECKLIST.md Makefile .gitkeep

# ---------------------------------------------------
# Version Control Operations
# ---------------------------------------------------

.PHONY: git-commit-push
git-commit-push:
	@read -p "Enter commit message: " commit_message; \
	if [ -z "$$commit_message" ]; then \
		echo "❌ Error: Commit message cannot be empty"; \
		exit 1; \
	fi; \
	git add .; \
	git commit -m "$$commit_message"; \
	git push

.PHONY: git-tag
git-tag:
	@bash -c '\
	read -p "Tag type (major/minor/patch): " tag_type; \
	last_tag=$$(git tag --sort=-v:refname | head -n 1); \
	if [ -z "$$last_tag" ]; then last_tag="0.0.0"; fi; \
	major=$$(echo $$last_tag | cut -d. -f1); \
	minor=$$(echo $$last_tag | cut -d. -f2); \
	patch=$$(echo $$last_tag | cut -d. -f3); \
	case "$$tag_type" in \
		major) major=$$((major + 1)); minor=0; patch=0;; \
		minor) minor=$$((minor + 1)); patch=0;; \
		patch) patch=$$((patch + 1));; \
		*) echo "❌ Invalid tag type: $$tag_type"; exit 1;; \
	esac; \
	new_tag="$$major.$$minor.$$patch"; \
	git tag -a "$$new_tag" -m "Release $$new_tag"; \
	git push origin "$$new_tag"; \
	echo "✅ Released new tag: $$new_tag"; \
	'

.PHONY: generate-ai-diff
generate-ai-diff:
	@mkdir -p diff
	@timestamp=$$(date +"%Y%m%d_%H%M%S"); \
	read -p "📁 Enter directory/path(s) to include in the diff (space-separated, leave empty for all changes): " DIR_PATHS; \
	if [ -z "$$DIR_PATHS" ]; then \
		echo "📝 Generating git diff for ALL changes into diff/diff_$${timestamp}.txt..."; \
		echo "Tu es un expert en revue de code et en conventions de commits (Conventional Commits)." > diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "À partir du diff Git ci-dessous, fais les choses suivantes :" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "1. Propose un nom de commit clair et concis en anglais" >> diff/diff_$${timestamp}.txt; \
		echo "   avec le format <type>(<scope>): <description>," >> diff/diff_$${timestamp}.txt; \
		echo "   en respectant les Conventional Commits" >> diff/diff_$${timestamp}.txt; \
		echo "   (ex: feat:, fix:, refactor:, test:, chore:, docs:)." >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "2. Rédige un résumé du travail effectué en quelques phrases," >> diff/diff_$${timestamp}.txt; \
		echo "   orienté métier et technique." >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "3. Donne une liste d'exemples concrets de changements, en t'appuyant sur le diff :" >> diff/diff_$${timestamp}.txt; \
		echo "   - méthodes ajoutées, modifiées ou supprimées" >> diff/diff_$${timestamp}.txt; \
		echo "   - responsabilités déplacées ou clarifiées" >> diff/diff_$${timestamp}.txt; \
		echo "   - améliorations de validation, de logique ou de structure" >> diff/diff_$${timestamp}.txt; \
		echo "   - impacts fonctionnels éventuels" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "Contraintes :" >> diff/diff_$${timestamp}.txt; \
		echo "   - Ne décris que ce qui est réellement visible dans le diff" >> diff/diff_$${timestamp}.txt; \
		echo "   - Sois précis, factuel et structuré" >> diff/diff_$${timestamp}.txt; \
		echo "   - Évite les suppositions" >> diff/diff_$${timestamp}.txt; \
		echo "   - Utilise un ton professionnel" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "4. SI et SEULEMENT SI les changements sont cassants (breaking changes) :" >> diff/diff_$${timestamp}.txt; \
		echo "   - Génère une entrée de CHANGELOG conforme à Keep a Changelog et SemVer." >> diff/diff_$${timestamp}.txt; \
		echo "   - Le changelog doit apparaître APRES les recommandations ci-dessus." >> diff/diff_$${timestamp}.txt; \
		echo "   - Utilise STRICTEMENT la structure suivante :" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "     ## [X.0.0] - YYYY-MM-DD" >> diff/diff_$${timestamp}.txt; \
		echo "     ### Changed" >> diff/diff_$${timestamp}.txt; \
		echo "     - Description claire du changement cassant" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "     ### Removed (si applicable)" >> diff/diff_$${timestamp}.txt; \
		echo "     - API, méthode ou comportement supprimé" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "     ### Security (si applicable)" >> diff/diff_$${timestamp}.txt; \
		echo "     - Impact sécurité lié au changement" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "   - Ne génère PAS de changelog si aucun breaking change n'est détecté." >> diff/diff_$${timestamp}.txt; \
		echo "   - N'invente PAS de version." >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "Voici le diff :" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		git diff HEAD -- . ':!*.phpunit.result.cache' ':!diff/*' >> diff/diff_$${timestamp}.txt; \
		echo "✅ Clean diff generated successfully: diff/diff_$${timestamp}.txt"; \
	else \
		echo "📝 Generating clean git diff for paths: $${DIR_PATHS} into diff/diff_$${timestamp}.txt..."; \
		echo "Tu es un expert en revue de code et en conventions de commits (Conventional Commits)." > diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "À partir du diff Git ci-dessous, fais les choses suivantes :" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "1. Propose un nom de commit clair et concis en anglais" >> diff/diff_$${timestamp}.txt; \
		echo "   avec le format <type>(<scope>): <description>," >> diff/diff_$${timestamp}.txt; \
		echo "   en respectant les Conventional Commits" >> diff/diff_$${timestamp}.txt; \
		echo "   (ex: feat:, fix:, refactor:, test:, chore:, docs:)." >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "2. Rédige un résumé du travail effectué en quelques phrases," >> diff/diff_$${timestamp}.txt; \
		echo "   orienté métier et technique." >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "3. Donne une liste d'exemples concrets de changements, en t'appuyant sur le diff :" >> diff/diff_$${timestamp}.txt; \
		echo "   - méthodes ajoutées, modifiées ou supprimées" >> diff/diff_$${timestamp}.txt; \
		echo "   - responsabilités déplacées ou clarifiées" >> diff/diff_$${timestamp}.txt; \
		echo "   - améliorations de validation, de logique ou de structure" >> diff/diff_$${timestamp}.txt; \
		echo "   - impacts fonctionnels éventuels" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "Contraintes :" >> diff/diff_$${timestamp}.txt; \
		echo "   - Ne décris que ce qui est réellement visible dans le diff" >> diff/diff_$${timestamp}.txt; \
		echo "   - Sois précis, factuel et structuré" >> diff/diff_$${timestamp}.txt; \
		echo "   - Évite les suppositions" >> diff/diff_$${timestamp}.txt; \
		echo "   - Utilise un ton professionnel" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "4. SI et SEULEMENT SI les changements sont cassants (breaking changes) :" >> diff/diff_$${timestamp}.txt; \
		echo "   - Génère une entrée de CHANGELOG conforme à Keep a Changelog et SemVer." >> diff/diff_$${timestamp}.txt; \
		echo "   - Le changelog doit apparaître APRES les recommandations ci-dessus." >> diff/diff_$${timestamp}.txt; \
		echo "   - Utilise STRICTEMENT la structure suivante :" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "     ## [X.0.0] - YYYY-MM-DD" >> diff/diff_$${timestamp}.txt; \
		echo "     ### Changed" >> diff/diff_$${timestamp}.txt; \
		echo "     - Description claire du changement cassant" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "     ### Removed (si applicable)" >> diff/diff_$${timestamp}.txt; \
		echo "     - API, méthode ou comportement supprimé" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "     ### Security (si applicable)" >> diff/diff_$${timestamp}.txt; \
		echo "     - Impact sécurité lié au changement" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "   - Ne génère PAS de changelog si aucun breaking change n'est détecté." >> diff/diff_$${timestamp}.txt; \
		echo "   - N'invente PAS de version." >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		echo "Voici le diff :" >> diff/diff_$${timestamp}.txt; \
		echo "" >> diff/diff_$${timestamp}.txt; \
		git diff HEAD -- $$DIR_PATHS ':!*.phpunit.result.cache' ':!diff/*' >> diff/diff_$${timestamp}.txt; \
		echo "✅ Clean diff generated successfully: diff/diff_$${timestamp}.txt"; \
	fi

.PHONY: list-diffs
list-diffs:
	@echo "📁 Available diff files:"
	@ls -la diff/diff_*.txt 2>/dev/null || echo "No diff files found"

.PHONY: git-tag-republish
git-tag-republish:
	@bash -c '\
	last_tag=$$(git tag --sort=-v:refname | head -n 1); \
	if [ -z "$$last_tag" ]; then echo "❌ No tags found!"; exit 1; fi; \
	echo "Republishing last tag: $$last_tag"; \
	git push origin "$$last_tag" --force; \
	echo "✅ Tag $$last_tag republished"; \
	'

# ---------------------------------------------------
# File Management Operations
# ---------------------------------------------------

.PHONY: concat-all
concat-all:
	@read -p "📁 Enter the source directory path to scan (leave empty for default './src ./database ./tests'): " SOURCE_PATH; \
	if [ -z "$$SOURCE_PATH" ]; then \
		SOURCE_DIRS="./src ./database ./config ./tests"; \
		echo "🔗 Concatenating all files from default directories: $${SOURCE_DIRS} into all.txt..."; \
	else \
		SOURCE_DIRS="$$SOURCE_PATH"; \
		echo "🔗 Concatenating all files from directory: $${SOURCE_DIRS} into all.txt..."; \
	fi; \
	find $${SOURCE_DIRS} -type f -exec sh -c 'echo ""; echo "// ==== {} ==="; echo ""; cat {}' \; > all.txt; \
	echo "✅ File all.txt generated successfully from: $${SOURCE_DIRS}"

# ---------------------------------------------------
# Release Management Workflow
# ---------------------------------------------------

.PHONY: release
release:
	@echo "🚀 Creating release..."
	@make git-tag
	@echo "✅ Release created successfully"

# ---------------------------------------------------
# Help & Documentation
# ---------------------------------------------------

.PHONY: help
help:
	@echo "📚 Available commands:"
	@echo ""
	@echo "🚀 Version Control:"
	@echo "  git-commit-push       Commit and push all changes"
	@echo "  git-tag               Create and push a new version tag"
	@echo "  generate-ai-diff      Generate clean diff for AI review"
	@echo "  git-tag-republish     Force push the last tag"
	@echo ""
	@echo "📁 File Management:"
	@echo "  concat-all            Concatenate all files from directories"
	@echo ""
	@echo "🔄 Release Management:"
	@echo "  release               Create new release (includes pre-release)"
	@echo ""
	@echo "❓ Help:"
	@echo "  help                  Display this help message"

# ---------------------------------------------------
# Default Target
# ---------------------------------------------------
.DEFAULT_GOAL := help
