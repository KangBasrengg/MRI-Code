package analyzer

import (
	"fmt"
	"math"
	"strings"
)

// CalculateHealthScore synthesizes structural metrics from Dead Code detection, Circular Dependencies,
// documentation density, and symbol complexity to compute an authoritative grade from 0 to 100.
func CalculateHealthScore(
	totalNodes int,
	deadCodeCount int,
	circularCount int,
	totalLOC int,
	totalComments int,
	complexityMetrics []ComplexityMetric,
) HealthScore {
	score := 100.0
	var suggestions []string

	// 1. Documentation Density Analysis
	commentRatio := 0.0
	if totalLOC > 0 {
		commentRatio = math.Round((float64(totalComments)/float64(totalLOC))*1000) / 10
		if commentRatio < 6.0 {
			penalty := math.Min(8.0, (6.0-commentRatio)*1.5)
			score -= penalty
			suggestions = append(suggestions, fmt.Sprintf("📝 Documentation Gap: Current comment ratio is low (%.1f%%). Consider enriching complex functions with Go/TS docstrings for team velocity.", commentRatio))
		}
	}

	// 2. Dead Code Penalty
	deadRate := 0.0
	if totalNodes > 0 {
		deadRate = math.Round((float64(deadCodeCount)/float64(totalNodes))*1000) / 10
		if deadCodeCount > 0 {
			penalty := math.Min(20.0, float64(deadCodeCount)*1.2)
			score -= penalty
			suggestions = append(suggestions, fmt.Sprintf("💡 Clean Code & Footprint: Detected %d isolated symbols with zero incoming calls (%.1f%% dead code rate). Review and archive unused boilerplate.", deadCodeCount, deadRate))
		}
	}

	// 3. Circular Dependency Anti-Pattern Penalty (Heavy Deduction)
	if circularCount > 0 {
		penalty := math.Min(35.0, float64(circularCount)*12.0)
		score -= penalty
		suggestions = append(suggestions, fmt.Sprintf("🚨 Critical Architectural Debt: Discovered %d circular package dependency loop(s). Decouple cyclical bindings immediately by introducing abstraction interfaces or dedicated domain models.", circularCount))
	}

	// 4. Complexity & Hotspot Analysis
	extremeHotspots := 0
	highHotspots := 0
	for _, m := range complexityMetrics {
		if strings.Contains(m.Rating, "Extreme") {
			extremeHotspots++
		} else if strings.Contains(m.Rating, "High") {
			highHotspots++
		}
	}
	if extremeHotspots > 0 || highHotspots > 0 {
		penalty := math.Min(25.0, float64(extremeHotspots)*4.5+float64(highHotspots)*1.5)
		score -= penalty
		if extremeHotspots > 0 {
			suggestions = append(suggestions, fmt.Sprintf("⚙️ Complexity Hotspot Warning: Found %d file(s) labeled as Extreme structural debt. Consider modularizing monolithic files into focused atomic services.", extremeHotspots))
		}
	}

	// Ensure boundary clamping between 0 and 100
	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}
	finalScore := int(math.Round(score))

	// Determine Grade and Debt Status
	grade, status := determineGradeAndStatus(finalScore, circularCount)
	if len(suggestions) == 0 {
		suggestions = append(suggestions, "✨ Immaculate Repository Architecture: No critical dead code, no circular dependency bindings, and clean structural modularity detected!")
	}

	return HealthScore{
		OverallScore:  finalScore,
		Grade:         grade,
		DebtStatus:    status,
		CommentRatio:  commentRatio,
		DeadCodeRate:  deadRate,
		CircularCount: circularCount,
		Suggestions:   suggestions,
	}
}

func determineGradeAndStatus(score int, circularCount int) (string, string) {
	// A circular dependency prevents achieving an A+ regardless of score
	if score >= 94 && circularCount == 0 {
		return "A+", "Immaculate Architecture (Negligible Debt)"
	} else if score >= 87 && circularCount == 0 {
		return "A", "Robust System Structure (Low Technical Debt)"
	} else if score >= 78 {
		return "B", "Healthy Quality (Moderate Structural Debt)"
	} else if score >= 65 {
		return "C", "Noticeable Friction (Actionable Refactoring Needed)"
	} else if score >= 50 {
		return "D", "High Maintenance Risk (Tight Coupling Detected)"
	}
	return "F", "Critical Technical Debt (Immediate Overhaul Required)"
}
