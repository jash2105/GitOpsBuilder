

package v1alpha1

const (
	// HealthyCondition represents the last recorded
	// health assessment result.
	HealthyCondition string = "Healthy"

	// ArtifactFailedReason represents the fact that the
	// source artifact download failed.
	ArtifactFailedReason string = "ArtifactFailed"

	// BuildFailedReason represents the fact that the
	// docker build failed.
	BuildFailedReason string = "BuildFailed"

	// PushFailedReason represents the fact that the
	// docker push failed.
	PushFailedReason string = "PushFailed"
)
