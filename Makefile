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

.PHONY: pre-commit
pre-commit:
	@echo "🔍 Running pre-commit checks..."
	@rm -f all.txt diff.txt
	@echo "✅ Pre-commit checks passed"

.PHONY: git-commit-push
git-commit-push: pre-commit update-checklist
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

.PHONY: update-checklist
update-checklist:
	@echo "📋 Updating FILES_CHECKLIST.md..."
	@if [ -f FILES_CHECKLIST.md ]; then \
		grep -E '^[0-9]+\. .* \[[ xX]\]$$' FILES_CHECKLIST.md > .existing_checklist.tmp; \
		awk -F' ' '{ \
			file_path=""; \
			for(i=2;i<NF;i++) { \
				if(i>2) file_path=file_path" "; \
				file_path=file_path$$i; \
			} \
			checkmark_state=$$NF; \
			print file_path " " checkmark_state \
		}' .existing_checklist.tmp > .existing_files.tmp; \
	else \
		touch .existing_files.tmp; \
		touch FILES_CHECKLIST.md; \
	fi; \
	echo "# Project File Checklist" > FILES_CHECKLIST.md; \
	echo "*Last updated: $$(date)*" >> FILES_CHECKLIST.md; \
	echo "" >> FILES_CHECKLIST.md; \
	echo "## Previously Checked Files" >> FILES_CHECKLIST.md; \
	file_count=1; \
	grep '\[x\]' .existing_files.tmp | sort | uniq | while read -r line; do \
		file_path=$$(echo "$$line" | awk '{$$NF=""; print $$0}' | sed 's/ $$//'); \
		echo "$$file_count. $$file_path [x]" >> FILES_CHECKLIST.md; \
		file_count=$$((file_count + 1)); \
	done; \
	previously_checked_files=$$(grep '\[x\]' .existing_files.tmp | awk '{$$NF=""; print $$0}' | sed 's/ $$//'); \
	echo "" >> FILES_CHECKLIST.md; \
	echo "## Other Files" >> FILES_CHECKLIST.md; \
	file_count=1; \
	find $(SOURCE_DIRS) -type f | sort | while read -r file_path; do \
		if ! echo "$$previously_checked_files" | grep -Fxq "$$file_path" 2>/dev/null; then \
			echo "$$file_count. $$file_path [ ]" >> FILES_CHECKLIST.md; \
			file_count=$$((file_count + 1)); \
		fi; \
	done; \
	rm -f .existing_checklist.tmp .existing_files.tmp; \
	echo "✅ FILES_CHECKLIST.md updated successfully"

.PHONY: list-modified-files
list-modified-files:
	@echo "📝 Updating CHANGED_FILES.md..."
	@previously_checked_files=$$(grep -E '^[0-9]+\. .* \[[xX]\]' FILES_CHECKLIST.md | sed 's/^[0-9]\+\. //' | sed 's/ *\[[xX]\]$$//'); \
	modified_file_count=0; \
	all_files=$$( (git diff --name-only; git ls-files --others --exclude-standard) | sort -u ); \
	echo "# Changed and Untracked Files" > CHANGED_FILES.md; \
	echo "*Updated: $$(date)*" >> CHANGED_FILES.md; \
	echo "" >> CHANGED_FILES.md; \
	echo "## Files to Review (modifications on checked files)" >> CHANGED_FILES.md; \
	for file_path in $$all_files; do \
		if echo "$$previously_checked_files" | grep -Fxq "$$file_path"; then \
			modified_file_count=$$((modified_file_count + 1)); \
			echo "$$modified_file_count. $$file_path [x]" >> CHANGED_FILES.md; \
		fi; \
	done; \
	if [ $$modified_file_count -eq 0 ]; then \
		echo "*(No modified files in this category)*" >> CHANGED_FILES.md; \
	fi; \
	echo "" >> CHANGED_FILES.md; \
	echo "## Other Modified Files" >> CHANGED_FILES.md; \
	modified_file_count=0; \
	for file_path in $$all_files; do \
		should_skip_file=0; \
		for ignored_file in $$(echo -e "$(IGNORED_FILES)"); do \
			if [ "$$file_path" = "$$ignored_file" ]; then should_skip_file=1; break; fi; \
		done; \
		if [ $$should_skip_file -eq 0 ] && ! echo "$$previously_checked_files" | grep -Fxq "$$file_path"; then \
			modified_file_count=$$((modified_file_count + 1)); \
			echo "$$modified_file_count. $$file_path [ ]" >> CHANGED_FILES.md; \
		fi; \
	done; \
	if [ $$modified_file_count -eq 0 ]; then \
		echo "*(No modified files in this category)*" >> CHANGED_FILES.md; \
	fi; \
	echo "✅ CHANGED_FILES.md updated successfully"

.PHONY: update-all
update-all: update-checklist list-modified-files
	@echo "✅ All file management updates completed"

# ---------------------------------------------------
# Release Management Workflow
# ---------------------------------------------------

.PHONY: pre-release
pre-release:
	@echo "🚀 Running pre-release checks..."
	@echo "✅ Pre-release checks completed"

.PHONY: release
release: pre-release
	@echo "🚀 Creating release..."
	@make git-tag
	@echo "✅ Release created successfully"

.PHONY: post-release
post-release:
	@echo "🧹 Performing post-release cleanup..."
	@make update-all
	@echo "✅ Post-release cleanup completed"

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
	@echo "  update-checklist      Update file checklist"
	@echo "  list-modified-files   List modified files"
	@echo "  update-all            Update checklist and modified files"
	@echo "  concat-all            Concatenate all files from directories"
	@echo ""
	@echo "🔄 Release Management:"
	@echo "  pre-release           Run all pre-release checks"
	@echo "  release               Create new release (includes pre-release)"
	@echo "  post-release          Clean up after release"
	@echo ""
	@echo "❓ Help:"
	@echo "  help                  Display this help message"

# ---------------------------------------------------
# Default Target
# ---------------------------------------------------
.DEFAULT_GOAL := help
