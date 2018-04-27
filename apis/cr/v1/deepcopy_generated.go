// +build !ignore_autogenerated

// crunchy

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgContainerResources) DeepCopyInto(out *PgContainerResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgContainerResources.
func (in *PgContainerResources) DeepCopy() *PgContainerResources {
	if in == nil {
		return nil
	}
	out := new(PgContainerResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgStorageSpec) DeepCopyInto(out *PgStorageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgStorageSpec.
func (in *PgStorageSpec) DeepCopy() *PgStorageSpec {
	if in == nil {
		return nil
	}
	out := new(PgStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgbackup) DeepCopyInto(out *Pgbackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgbackup.
func (in *Pgbackup) DeepCopy() *Pgbackup {
	if in == nil {
		return nil
	}
	out := new(Pgbackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgbackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgbackupList) DeepCopyInto(out *PgbackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgbackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgbackupList.
func (in *PgbackupList) DeepCopy() *PgbackupList {
	if in == nil {
		return nil
	}
	out := new(PgbackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgbackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgbackupSpec) DeepCopyInto(out *PgbackupSpec) {
	*out = *in
	out.StorageSpec = in.StorageSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgbackupSpec.
func (in *PgbackupSpec) DeepCopy() *PgbackupSpec {
	if in == nil {
		return nil
	}
	out := new(PgbackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgbackupStatus) DeepCopyInto(out *PgbackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgbackupStatus.
func (in *PgbackupStatus) DeepCopy() *PgbackupStatus {
	if in == nil {
		return nil
	}
	out := new(PgbackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgcluster) DeepCopyInto(out *Pgcluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgcluster.
func (in *Pgcluster) DeepCopy() *Pgcluster {
	if in == nil {
		return nil
	}
	out := new(Pgcluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgcluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgclusterList) DeepCopyInto(out *PgclusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgcluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgclusterList.
func (in *PgclusterList) DeepCopy() *PgclusterList {
	if in == nil {
		return nil
	}
	out := new(PgclusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgclusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgclusterSpec) DeepCopyInto(out *PgclusterSpec) {
	*out = *in
	out.PrimaryStorage = in.PrimaryStorage
	out.ReplicaStorage = in.ReplicaStorage
	out.ContainerResources = in.ContainerResources
	if in.UserLabels != nil {
		in, out := &in.UserLabels, &out.UserLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgclusterSpec.
func (in *PgclusterSpec) DeepCopy() *PgclusterSpec {
	if in == nil {
		return nil
	}
	out := new(PgclusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgclusterStatus) DeepCopyInto(out *PgclusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgclusterStatus.
func (in *PgclusterStatus) DeepCopy() *PgclusterStatus {
	if in == nil {
		return nil
	}
	out := new(PgclusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgingest) DeepCopyInto(out *Pgingest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgingest.
func (in *Pgingest) DeepCopy() *Pgingest {
	if in == nil {
		return nil
	}
	out := new(Pgingest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgingest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgingestList) DeepCopyInto(out *PgingestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgingest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgingestList.
func (in *PgingestList) DeepCopy() *PgingestList {
	if in == nil {
		return nil
	}
	out := new(PgingestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgingestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgingestSpec) DeepCopyInto(out *PgingestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgingestSpec.
func (in *PgingestSpec) DeepCopy() *PgingestSpec {
	if in == nil {
		return nil
	}
	out := new(PgingestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgingestStatus) DeepCopyInto(out *PgingestStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgingestStatus.
func (in *PgingestStatus) DeepCopy() *PgingestStatus {
	if in == nil {
		return nil
	}
	out := new(PgingestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgpolicy) DeepCopyInto(out *Pgpolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgpolicy.
func (in *Pgpolicy) DeepCopy() *Pgpolicy {
	if in == nil {
		return nil
	}
	out := new(Pgpolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgpolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgpolicyList) DeepCopyInto(out *PgpolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgpolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgpolicyList.
func (in *PgpolicyList) DeepCopy() *PgpolicyList {
	if in == nil {
		return nil
	}
	out := new(PgpolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgpolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgpolicySpec) DeepCopyInto(out *PgpolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgpolicySpec.
func (in *PgpolicySpec) DeepCopy() *PgpolicySpec {
	if in == nil {
		return nil
	}
	out := new(PgpolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgpolicyStatus) DeepCopyInto(out *PgpolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgpolicyStatus.
func (in *PgpolicyStatus) DeepCopy() *PgpolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PgpolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgreplica) DeepCopyInto(out *Pgreplica) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgreplica.
func (in *Pgreplica) DeepCopy() *Pgreplica {
	if in == nil {
		return nil
	}
	out := new(Pgreplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgreplica) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgreplicaList) DeepCopyInto(out *PgreplicaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgreplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgreplicaList.
func (in *PgreplicaList) DeepCopy() *PgreplicaList {
	if in == nil {
		return nil
	}
	out := new(PgreplicaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgreplicaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgreplicaSpec) DeepCopyInto(out *PgreplicaSpec) {
	*out = *in
	out.ReplicaStorage = in.ReplicaStorage
	out.ContainerResources = in.ContainerResources
	if in.UserLabels != nil {
		in, out := &in.UserLabels, &out.UserLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgreplicaSpec.
func (in *PgreplicaSpec) DeepCopy() *PgreplicaSpec {
	if in == nil {
		return nil
	}
	out := new(PgreplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgreplicaStatus) DeepCopyInto(out *PgreplicaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgreplicaStatus.
func (in *PgreplicaStatus) DeepCopy() *PgreplicaStatus {
	if in == nil {
		return nil
	}
	out := new(PgreplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgtask) DeepCopyInto(out *Pgtask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgtask.
func (in *Pgtask) DeepCopy() *Pgtask {
	if in == nil {
		return nil
	}
	out := new(Pgtask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgtask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgtaskList) DeepCopyInto(out *PgtaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgtask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgtaskList.
func (in *PgtaskList) DeepCopy() *PgtaskList {
	if in == nil {
		return nil
	}
	out := new(PgtaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgtaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgtaskSpec) DeepCopyInto(out *PgtaskSpec) {
	*out = *in
	out.StorageSpec = in.StorageSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgtaskSpec.
func (in *PgtaskSpec) DeepCopy() *PgtaskSpec {
	if in == nil {
		return nil
	}
	out := new(PgtaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgtaskStatus) DeepCopyInto(out *PgtaskStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgtaskStatus.
func (in *PgtaskStatus) DeepCopy() *PgtaskStatus {
	if in == nil {
		return nil
	}
	out := new(PgtaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pgupgrade) DeepCopyInto(out *Pgupgrade) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pgupgrade.
func (in *Pgupgrade) DeepCopy() *Pgupgrade {
	if in == nil {
		return nil
	}
	out := new(Pgupgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pgupgrade) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgupgradeList) DeepCopyInto(out *PgupgradeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pgupgrade, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgupgradeList.
func (in *PgupgradeList) DeepCopy() *PgupgradeList {
	if in == nil {
		return nil
	}
	out := new(PgupgradeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgupgradeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgupgradeSpec) DeepCopyInto(out *PgupgradeSpec) {
	*out = *in
	out.StorageSpec = in.StorageSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgupgradeSpec.
func (in *PgupgradeSpec) DeepCopy() *PgupgradeSpec {
	if in == nil {
		return nil
	}
	out := new(PgupgradeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgupgradeStatus) DeepCopyInto(out *PgupgradeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgupgradeStatus.
func (in *PgupgradeStatus) DeepCopy() *PgupgradeStatus {
	if in == nil {
		return nil
	}
	out := new(PgupgradeStatus)
	in.DeepCopyInto(out)
	return out
}
