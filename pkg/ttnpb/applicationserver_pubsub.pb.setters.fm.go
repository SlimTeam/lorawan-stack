// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	time "time"

	types "github.com/gogo/protobuf/types"
)

func (dst *ApplicationPubSubIdentifiers) SetFields(src *ApplicationPubSubIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				newDst = &dst.ApplicationIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "pub_sub_id":
			if len(subs) > 0 {
				return fmt.Errorf("'pub_sub_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PubSubID = src.PubSubID
			} else {
				var zero string
				dst.PubSubID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPubSub) SetFields(src *ApplicationPubSub, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSubIdentifiers
				if src != nil {
					newSrc = &src.ApplicationPubSubIdentifiers
				}
				newDst = &dst.ApplicationPubSubIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationPubSubIdentifiers = src.ApplicationPubSubIdentifiers
				} else {
					var zero ApplicationPubSubIdentifiers
					dst.ApplicationPubSubIdentifiers = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "format":
			if len(subs) > 0 {
				return fmt.Errorf("'format' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Format = src.Format
			} else {
				var zero string
				dst.Format = zero
			}
		case "base_topic":
			if len(subs) > 0 {
				return fmt.Errorf("'base_topic' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BaseTopic = src.BaseTopic
			} else {
				var zero string
				dst.BaseTopic = zero
			}
		case "downlink_push":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.DownlinkPush == nil) && dst.DownlinkPush == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkPush
				}
				if dst.DownlinkPush != nil {
					newDst = dst.DownlinkPush
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.DownlinkPush = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkPush = src.DownlinkPush
				} else {
					dst.DownlinkPush = nil
				}
			}
		case "downlink_replace":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.DownlinkReplace == nil) && dst.DownlinkReplace == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkReplace
				}
				if dst.DownlinkReplace != nil {
					newDst = dst.DownlinkReplace
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.DownlinkReplace = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkReplace = src.DownlinkReplace
				} else {
					dst.DownlinkReplace = nil
				}
			}
		case "uplink_message":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.UplinkMessage == nil) && dst.UplinkMessage == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkMessage
				}
				if dst.UplinkMessage != nil {
					newDst = dst.UplinkMessage
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.UplinkMessage = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkMessage = src.UplinkMessage
				} else {
					dst.UplinkMessage = nil
				}
			}
		case "join_accept":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.JoinAccept == nil) && dst.JoinAccept == nil {
					continue
				}
				if src != nil {
					newSrc = src.JoinAccept
				}
				if dst.JoinAccept != nil {
					newDst = dst.JoinAccept
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.JoinAccept = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.JoinAccept = src.JoinAccept
				} else {
					dst.JoinAccept = nil
				}
			}
		case "downlink_ack":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.DownlinkAck == nil) && dst.DownlinkAck == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkAck
				}
				if dst.DownlinkAck != nil {
					newDst = dst.DownlinkAck
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.DownlinkAck = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkAck = src.DownlinkAck
				} else {
					dst.DownlinkAck = nil
				}
			}
		case "downlink_nack":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.DownlinkNack == nil) && dst.DownlinkNack == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkNack
				}
				if dst.DownlinkNack != nil {
					newDst = dst.DownlinkNack
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.DownlinkNack = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkNack = src.DownlinkNack
				} else {
					dst.DownlinkNack = nil
				}
			}
		case "downlink_sent":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.DownlinkSent == nil) && dst.DownlinkSent == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkSent
				}
				if dst.DownlinkSent != nil {
					newDst = dst.DownlinkSent
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.DownlinkSent = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkSent = src.DownlinkSent
				} else {
					dst.DownlinkSent = nil
				}
			}
		case "downlink_failed":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.DownlinkFailed == nil) && dst.DownlinkFailed == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkFailed
				}
				if dst.DownlinkFailed != nil {
					newDst = dst.DownlinkFailed
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.DownlinkFailed = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkFailed = src.DownlinkFailed
				} else {
					dst.DownlinkFailed = nil
				}
			}
		case "downlink_queued":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.DownlinkQueued == nil) && dst.DownlinkQueued == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkQueued
				}
				if dst.DownlinkQueued != nil {
					newDst = dst.DownlinkQueued
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.DownlinkQueued = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkQueued = src.DownlinkQueued
				} else {
					dst.DownlinkQueued = nil
				}
			}
		case "location_solved":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.LocationSolved == nil) && dst.LocationSolved == nil {
					continue
				}
				if src != nil {
					newSrc = src.LocationSolved
				}
				if dst.LocationSolved != nil {
					newDst = dst.LocationSolved
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.LocationSolved = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.LocationSolved = src.LocationSolved
				} else {
					dst.LocationSolved = nil
				}
			}
		case "service_data":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub_Message
				if (src == nil || src.ServiceData == nil) && dst.ServiceData == nil {
					continue
				}
				if src != nil {
					newSrc = src.ServiceData
				}
				if dst.ServiceData != nil {
					newDst = dst.ServiceData
				} else {
					newDst = &ApplicationPubSub_Message{}
					dst.ServiceData = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ServiceData = src.ServiceData
				} else {
					dst.ServiceData = nil
				}
			}

		case "provider":
			if len(subs) == 0 && src == nil {
				dst.Provider = nil
				continue
			} else if len(subs) == 0 {
				dst.Provider = src.Provider
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "nats":
					_, srcOk := src.Provider.(*ApplicationPubSub_NATS)
					if !srcOk && src.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'nats', while different oneof is set in source")
					}
					_, dstOk := dst.Provider.(*ApplicationPubSub_NATS)
					if !dstOk && dst.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'nats', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ApplicationPubSub_NATSProvider
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Provider.(*ApplicationPubSub_NATS).NATS
						}
						if dstOk {
							newDst = dst.Provider.(*ApplicationPubSub_NATS).NATS
						} else {
							newDst = &ApplicationPubSub_NATSProvider{}
							dst.Provider = &ApplicationPubSub_NATS{NATS: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Provider = src.Provider
						} else {
							dst.Provider = nil
						}
					}
				case "mqtt":
					_, srcOk := src.Provider.(*ApplicationPubSub_MQTT)
					if !srcOk && src.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'mqtt', while different oneof is set in source")
					}
					_, dstOk := dst.Provider.(*ApplicationPubSub_MQTT)
					if !dstOk && dst.Provider != nil {
						return fmt.Errorf("attempt to set oneof 'mqtt', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ApplicationPubSub_MQTTProvider
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Provider.(*ApplicationPubSub_MQTT).MQTT
						}
						if dstOk {
							newDst = dst.Provider.(*ApplicationPubSub_MQTT).MQTT
						} else {
							newDst = &ApplicationPubSub_MQTTProvider{}
							dst.Provider = &ApplicationPubSub_MQTT{MQTT: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Provider = src.Provider
						} else {
							dst.Provider = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPubSubs) SetFields(src *ApplicationPubSubs, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "pubsubs":
			if len(subs) > 0 {
				return fmt.Errorf("'pubsubs' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Pubsubs = src.Pubsubs
			} else {
				dst.Pubsubs = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPubSubFormats) SetFields(src *ApplicationPubSubFormats, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "formats":
			if len(subs) > 0 {
				return fmt.Errorf("'formats' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formats = src.Formats
			} else {
				dst.Formats = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetApplicationPubSubRequest) SetFields(src *GetApplicationPubSubRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSubIdentifiers
				if src != nil {
					newSrc = &src.ApplicationPubSubIdentifiers
				}
				newDst = &dst.ApplicationPubSubIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationPubSubIdentifiers = src.ApplicationPubSubIdentifiers
				} else {
					var zero ApplicationPubSubIdentifiers
					dst.ApplicationPubSubIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListApplicationPubSubsRequest) SetFields(src *ListApplicationPubSubsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				newDst = &dst.ApplicationIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetApplicationPubSubRequest) SetFields(src *SetApplicationPubSubRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "pubsub":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPubSub
				if src != nil {
					newSrc = &src.ApplicationPubSub
				}
				newDst = &dst.ApplicationPubSub
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationPubSub = src.ApplicationPubSub
				} else {
					var zero ApplicationPubSub
					dst.ApplicationPubSub = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPubSub_NATSProvider) SetFields(src *ApplicationPubSub_NATSProvider, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "server_url":
			if len(subs) > 0 {
				return fmt.Errorf("'server_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ServerURL = src.ServerURL
			} else {
				var zero string
				dst.ServerURL = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPubSub_MQTTProvider) SetFields(src *ApplicationPubSub_MQTTProvider, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "server_url":
			if len(subs) > 0 {
				return fmt.Errorf("'server_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ServerURL = src.ServerURL
			} else {
				var zero string
				dst.ServerURL = zero
			}
		case "client_id":
			if len(subs) > 0 {
				return fmt.Errorf("'client_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ClientID = src.ClientID
			} else {
				var zero string
				dst.ClientID = zero
			}
		case "username":
			if len(subs) > 0 {
				return fmt.Errorf("'username' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Username = src.Username
			} else {
				var zero string
				dst.Username = zero
			}
		case "password":
			if len(subs) > 0 {
				return fmt.Errorf("'password' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Password = src.Password
			} else {
				var zero string
				dst.Password = zero
			}
		case "subscribe_qos":
			if len(subs) > 0 {
				return fmt.Errorf("'subscribe_qos' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SubscribeQoS = src.SubscribeQoS
			} else {
				var zero ApplicationPubSub_MQTTProvider_QoS
				dst.SubscribeQoS = zero
			}
		case "publish_qos":
			if len(subs) > 0 {
				return fmt.Errorf("'publish_qos' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PublishQoS = src.PublishQoS
			} else {
				var zero ApplicationPubSub_MQTTProvider_QoS
				dst.PublishQoS = zero
			}
		case "use_tls":
			if len(subs) > 0 {
				return fmt.Errorf("'use_tls' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UseTLS = src.UseTLS
			} else {
				var zero bool
				dst.UseTLS = zero
			}
		case "tls_ca":
			if len(subs) > 0 {
				return fmt.Errorf("'tls_ca' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TLSCA = src.TLSCA
			} else {
				dst.TLSCA = nil
			}
		case "tls_client_cert":
			if len(subs) > 0 {
				return fmt.Errorf("'tls_client_cert' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TLSClientCert = src.TLSClientCert
			} else {
				dst.TLSClientCert = nil
			}
		case "tls_client_key":
			if len(subs) > 0 {
				return fmt.Errorf("'tls_client_key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TLSClientKey = src.TLSClientKey
			} else {
				dst.TLSClientKey = nil
			}
		case "headers":
			if len(subs) > 0 {
				return fmt.Errorf("'headers' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Headers = src.Headers
			} else {
				dst.Headers = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPubSub_Message) SetFields(src *ApplicationPubSub_Message, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "topic":
			if len(subs) > 0 {
				return fmt.Errorf("'topic' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Topic = src.Topic
			} else {
				var zero string
				dst.Topic = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
